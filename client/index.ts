import { Anthropic } from "@anthropic-ai/sdk";
import {
  MessageParam as AnthropicMessage,
  Tool as AnthropicTool,
} from "@anthropic-ai/sdk/resources/messages/messages.mjs";

import { Client } from "@modelcontextprotocol/sdk/client/index.js";
import { StdioClientTransport } from "@modelcontextprotocol/sdk/client/stdio.js";
import readline from "readline/promises";
import dotenv from "dotenv";

dotenv.config();

const ANTHROPIC_API_KEY = process.env.ANTHROPIC_API_KEY;
const ANTHROPIC_MODEL = process.env.ANTHROPIC_MODEL;
const ANTHROPIC_MAX_TOKENS = process.env.ANTHROPIC_MAX_TOKENS || 1000;

// Client Structure
class MCPClient {
  private mcp: Client;
  private anthropic: Anthropic | null = null;
  private transport: StdioClientTransport | null = null;
  private anthropicTools: AnthropicTool[] = [];

  constructor() {
    this.anthropic = new Anthropic({
      apiKey: ANTHROPIC_API_KEY,
    });

    this.mcp = new Client({ name: "mcp-client-cli", version: "1.0.0" });
  }

  // Connection Management
  async connectToServer(serverScriptPath: string) {
    try {
      this.transport = new StdioClientTransport({
        command: serverScriptPath,
        args: [],
      });
      this.mcp.connect(this.transport);

      const toolsResult = await this.mcp.listTools();

      this.anthropicTools = toolsResult.tools.map((tool) => {
        return {
          name: tool.name,
          description: tool.description,
          input_schema: tool.inputSchema,
        };
      });
      console.log(
        "Connected to server with tools:",
        this.anthropicTools.map(({ name }) => name)
      );

    } catch (e) {
      console.log("Failed to connect to MCP server: ", e);
      throw e;
    }
  }

  // Query Processing Logic Anthropic
  async processQueryAnthropic(query: string) {
    const messages: AnthropicMessage[] = [
      {
        role: "user",
        content: query,
      },
    ];

    if (!this.anthropic) return;

    const response = await this.anthropic.messages.create({
      model: ANTHROPIC_MODEL!,
      max_tokens: +ANTHROPIC_MAX_TOKENS,
      messages,
      tools: this.anthropicTools,
    });

    const finalText = [];
    const toolResults = [];

    for (const content of response.content) {
      if (content.type === "text") {
        finalText.push(content.text);
      } else if (content.type === "tool_use") {
        const toolName = content.name;
        const toolArgs = content.input as { [x: string]: unknown } | undefined;

        const result = await this.mcp.callTool({
          name: toolName,
          arguments: toolArgs,
        });
        toolResults.push(result);
        finalText.push(
          `[Calling tool ${toolName} with args ${JSON.stringify(toolArgs)}]`
        );

        messages.push({
          role: "user",
          content: result.content as string,
        });

        const response = await this.anthropic.messages.create({
          model: ANTHROPIC_MODEL!,
          max_tokens: +ANTHROPIC_MAX_TOKENS,
          messages,
        });

        finalText.push(
          response.content[0].type === "text" ? response.content[0].text : ""
        );
      }
    }

    return finalText.join("\n");
  }

  // Interactive Chat interface
  async chatLoop() {
    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout,
    });

    try {
      console.log("\nMCP Client Started!");
      console.log("Type your queries or 'quit' to exit.");

      while (true) {
        const message = await rl.question("\nQuery: ");
        if (message.toLowerCase() === "quit") {
          break;
        }

        const response = await this.processQueryAnthropic(message)

        console.log("\n" + response);
      }
    } finally {
      rl.close();
    }
  }

  async cleanup() {
    await this.mcp.close();
  }
}

// Main Entry Point
async function main() {
  if (process.argv.length < 3) {
    console.log("Usage: node index.ts <path_to_server_script>");
    return;
  }
  const mcpClient = new MCPClient();
  try {
    await mcpClient.connectToServer(process.argv[2]);
    await mcpClient.chatLoop();
  } finally {
    await mcpClient.cleanup();
    process.exit(0);
  }
}

// Run
main();
