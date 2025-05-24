import { defineConfig } from 'vocs'

export default defineConfig({
  title: 'MCP-Go',
  baseUrl: 'https://mark3labs.github.io',
  basePath: '/mcp-go',
  logoUrl: '/logo.png',
  description: 'A Go implementation of the Model Context Protocol (MCP), enabling seamless integration between LLM applications and external data sources and tools.',
  sidebar: [
    {
      text: 'Getting Started',
      link: '/getting-started',
    },
    {
      text: 'Example',
      link: '/example',
    },
  ],
  socials: [
    {
      icon: 'github',
      link: 'https://github.com/mark3labs/mcp-go',
    },
  ],
})
