import { Avatar, Group, Skeleton, Text, UnstyledButton } from '@mantine/core';
import { User } from '@/types/user';
import classes from './user-button.module.css';

interface UserButtonProps {
  user: User | null;
}

export function UserButton({ user }: UserButtonProps) {
  return (
    <Skeleton visible={!user}>
      <UnstyledButton className={classes.user}>
        <Group>
          <Avatar src={user ? user.image : ''} radius="xl" />

          <div style={{ flex: 1 }}>
            <Text size="sm" fw={500}>
              {user ? user.name : ''}
            </Text>

            <Text c="dimmed" size="xs">
              {user ? user.username : ''}
            </Text>
          </div>
        </Group>
      </UnstyledButton>
    </Skeleton>
  );
}
