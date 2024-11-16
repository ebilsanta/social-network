'use client';

import { Avatar, Button, Center, Flex, Skeleton, Text } from '@mantine/core';
import { useProfileHeader } from '@/app/[username]/_components/ProfileHeader/useProfileHeader';
import Loading from '@/components/loading';
import NotFound from '@/components/not-found';
import { User } from '@/types/user';

interface Stat {
  field: keyof User;
  label: string;
}

const stats: Stat[] = [
  { field: 'postCount', label: 'Posts' },
  { field: 'followerCount', label: 'Followers' },
  { field: 'followingCount', label: 'Following' },
];

interface ProfileHeaderProps {
  username: string | undefined;
}

export const ProfileHeader = ({ username }: ProfileHeaderProps) => {
  const { data: user, error } = useProfileHeader(username);
  if (error) {
    return <NotFound />;
  }
  if (!user) {
    return <Loading />;
  }

  return (
    <Skeleton visible={!user} width={600}>
      <Flex py={32} justify="center">
        <Center style={{ marginRight: 40 }}>
          <Avatar src={user?.image} size={120} />
        </Center>
        <Flex direction="column" gap="lg" justify="center">
          <Flex justify="space-between" align="center">
            <Text fz="lg" fw={600}>
              {user?.username}
            </Text>
            <Button size="xs" variant="variant">
              Follow
            </Button>
          </Flex>
          <Flex gap="xl" align="center">
            {stats.map((stat) => (
              <Text key={stat.label} fz="sm" fw={600}>
                {`${user?.[stat.field]} ${stat.label}`}
              </Text>
            ))}
          </Flex>
          <Text fw={500}>{user?.name}</Text>
        </Flex>
      </Flex>
    </Skeleton>
  );
};
