import { redirect } from 'next/navigation';
import { Avatar, Group, Skeleton, Text, UnstyledButton } from '@mantine/core';
import { User } from '@/types/user';
import classes from './ProfileCard.module.css';

interface ProfileCardProps {
  user: User | null;
  isSmallScreen?: boolean;
}

export function ProfileCard({ user, isSmallScreen }: ProfileCardProps) {
  return (
    <Skeleton visible={!user}>
      <UnstyledButton
        className={classes.user}
        data-minimized={isSmallScreen || undefined}
        onClick={() => redirect(user?.username!)}
      >
        <Group justify="center">
          <Avatar src={user ? user.image : ''} radius="xl" />

          <div style={{ flex: 1 }}>
            <Text size="sm" fw={500}>
              {user && !isSmallScreen ? user.name : null}
            </Text>

            <Text c="dimmed" size="xs">
              {user && !isSmallScreen ? user.username : null}
            </Text>
          </div>
        </Group>
      </UnstyledButton>
    </Skeleton>
  );
}
