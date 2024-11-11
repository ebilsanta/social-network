import '@mantine/core/styles.css';

import React from 'react';
import { ColorSchemeScript, MantineProvider } from '@mantine/core';
import { SessionProvider } from '@/providers/next-auth-provider';
import { UserProvider } from '@/providers/user-provider';
import { theme } from '../theme';

export const metadata = {
  title: 'Mantine Next.js template',
  description: 'I am using Mantine with Next.js!',
};

export default function RootLayout({ children }: { children: any }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body
        style={{
          margin: 0,
          height: '100vh',
          display: 'flex',
          justifyContent: 'center',
        }}
      >
        <SessionProvider>
          <UserProvider>
            <MantineProvider theme={theme} defaultColorScheme="auto">
              <div style={{ width: '100%' }}>{children}</div>
            </MantineProvider>
          </UserProvider>
        </SessionProvider>
      </body>
    </html>
  );
}
