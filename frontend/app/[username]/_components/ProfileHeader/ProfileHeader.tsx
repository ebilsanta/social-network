'use client';

import { Avatar, Center, Flex, Skeleton, Text } from '@mantine/core';
import { FollowButton } from '@/app/[username]/_components/ProfileHeader/FollowButton/FollowButton';
import { useProfileHeader } from '@/app/[username]/_components/ProfileHeader/useProfileHeader';
import Loading from '@/components/Loading';
import NotFound from '@/components/NotFound';
import { User } from '@/types/user';

interface Stat {
  field: keyof User;
  label: string;
  pluralLabel: string;
}

const stats: Stat[] = [
  { field: 'postCount', label: 'Post', pluralLabel: 'Posts' },
  { field: 'followerCount', label: 'Follower', pluralLabel: 'Followers' },
  { field: 'followingCount', label: 'Following', pluralLabel: 'Following' },
];

interface ProfileHeaderProps {
  username: string | undefined;
}

export const ProfileHeader = ({ username }: ProfileHeaderProps) => {
  const {
    currentUser,
    profileUser,
    profileUserError,
    isFollowing,
    handleFollowUser,
    handleUnfollowUser,
  } = useProfileHeader(username);
  if (profileUserError) {
    return <NotFound />;
  }
  if (!profileUser) {
    return <Loading />;
  }

  return (
    <Center>
      <Skeleton visible={!profileUser}>
        <Flex py={32} justify="center">
          <Center style={{ marginRight: 40 }}>
            <Avatar src={profileUser?.image} size={120} />
          </Center>
          <Flex direction="column" gap="lg" justify="center">
            <Flex justify="space-between" align="center">
              <Text fz="lg" fw={600}>
                {profileUser?.username}
              </Text>
              <FollowButton
                currentUser={currentUser}
                profileUser={profileUser}
                isFollowing={isFollowing}
                handleFollowUser={handleFollowUser}
                handleUnfollowUser={handleUnfollowUser}
              />
            </Flex>
            <Flex gap="xl" align="center">
              {stats.map((stat) => (
                <Text key={stat.label} fz="sm" fw={600}>
                  {`${profileUser?.[stat.field]} ${profileUser?.[stat.field] !== 1 ? stat.pluralLabel : stat.label}`}
                </Text>
              ))}
            </Flex>
            <Text fw={500}>{profileUser?.name}</Text>
          </Flex>
        </Flex>
      </Skeleton>
    </Center>
  );
};
