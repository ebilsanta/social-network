'use client';

import { IconCircleCheck } from '@tabler/icons-react';
import { Button, Group, Modal } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { User } from '@/types/user';

interface FollowButtonProps {
  currentUser: User | null;
  profileUser: User | undefined;
  isFollowing: Boolean | undefined;
  handleFollowUser: () => void;
  handleUnfollowUser: () => void;
}

export const FollowButton = ({
  currentUser,
  profileUser,
  isFollowing,
  handleFollowUser,
  handleUnfollowUser,
}: FollowButtonProps) => {
  const [opened, { open, close }] = useDisclosure(false);

  if (!currentUser || !profileUser) {
    return null;
  }
  if (currentUser.id === profileUser.id) {
    return null;
  }
  return isFollowing ? (
    <>
      <Modal opened={opened} onClose={close} title="Confirm unfollow">
        Are you sure you want to unfollow {profileUser.username}?
        <Group mt="lg" justify="flex-end">
          <Button variant="default" onClick={close}>
            Cancel
          </Button>
          <Button
            onClick={() => {
              handleUnfollowUser();
              close();
            }}
            color="red"
          >
            Delete
          </Button>
        </Group>
      </Modal>
      <Button size="xs" variant="light" onClick={open}>
        Following &nbsp; <IconCircleCheck />
      </Button>
    </>
  ) : (
    <Button size="xs" variant="variant" onClick={handleFollowUser}>
      Follow
    </Button>
  );
};
