import '@mantine/core/styles.css';

import React from 'react';
import { ColorSchemeScript, MantineProvider } from '@mantine/core';
import { Layout } from '@/app/_components/Layout/Layout';
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
      <body>
        <SessionProvider>
          <UserProvider>
            <MantineProvider theme={theme} defaultColorScheme="auto">
              <Layout>{children}</Layout>
            </MantineProvider>
          </UserProvider>
        </SessionProvider>
      </body>
    </html>
  );
}
