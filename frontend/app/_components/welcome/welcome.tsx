'use client';

import { IconBrandGoogle } from '@tabler/icons-react';
import { signIn } from 'next-auth/react';
import { Button, Container, Group, Text } from '@mantine/core';
import classes from './welcome.module.css';

export function Welcome() {
  return (
    <div className={classes.wrapper}>
      <Container size={700} className={classes.inner}>
        <h1 className={classes.title}>
          Welcome to{' '}
          <Text component="span" variant="gradient" gradient={{ from: 'blue', to: 'cyan' }} inherit>
            Catchup
          </Text>
        </h1>

        <Text className={classes.description} c="dimmed">
          Follow your friends and stay up to date with their latest posts
        </Text>

        <Group className={classes.controls}>
          <Button
            size="xl"
            className={classes.control}
            variant="default"
            type="submit"
            leftSection={<IconBrandGoogle />}
            onClick={() => signIn('google')}
          >
            Sign in
          </Button>
        </Group>
      </Container>
    </div>
  );
}
