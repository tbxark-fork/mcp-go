import { defineConfig } from 'vocs'

export default defineConfig({
  title: 'MCP-Go',
  logoUrl: '/logo.png',
  description: 'A Go implementation of the Model Context Protocol (MCP), enabling seamless integration between LLM applications and external data sources and tools.',
  sidebar: [
    {
      text: 'Getting Started',
      link: '/getting-started',
    },
    {
      text: 'Quick Start',
      link: '/quick-start',
    },
    {
      text: 'Core Concepts',
      link: '/core-concepts',
    },
    {
      text: 'Building MCP Servers',
      collapsed: false,
      items: [
        {
          text: 'Overview',
          link: '/servers',
        },
        {
          text: 'Server Basics',
          link: '/servers/basics',
        },
        {
          text: 'Resources',
          link: '/servers/resources',
        },
        {
          text: 'Tools',
          link: '/servers/tools',
        },
        {
          text: 'Prompts',
          link: '/servers/prompts',
        },
        {
          text: 'Advanced Features',
          link: '/servers/advanced',
        },
      ],
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
