import useSWR, { mutate } from 'swr';
import { FollowerAPI } from '@/lib/follower-api';
import { UserAPI } from '@/lib/user-api';
import { useUser } from '@/providers/user-provider';

export const useProfileHeader = (username?: string) => {
  const { user: currentUser } = useUser();
  const fetchUser = async () => {
    const response = await UserAPI.getUserByUsername(username!);
    return response.data;
  };
  const checkFollowing = async () => {
    const response = await UserAPI.checkFollowing(currentUser?.id!, profileUser?.id!);
    return response.following;
  };
  const { data: profileUser, error: profileUserError } = useSWR(
    username ? `/api/users/username/${username}` : null,
    fetchUser
  );
  const { data: isFollowing, error: isFollowingError } = useSWR(
    profileUser && currentUser ? `/api/users/${currentUser.id}/following/${profileUser.id}` : null,
    checkFollowing
  );
  const handleFollowUser = async () => {
    await FollowerAPI.addFollower(currentUser!.id, profileUser!.id);
    mutate(`/api/users/${currentUser!.id}/following/${profileUser!.id}`);
  };
  const handleUnfollowUser = async () => {
    await FollowerAPI.deleteFollower(currentUser!.id, profileUser!.id);
    mutate(`/api/users/${currentUser!.id}/following/${profileUser!.id}`);
  };
  return {
    currentUser,
    profileUser,
    profileUserError,
    isFollowing,
    isFollowingError,
    handleFollowUser,
    handleUnfollowUser,
  };
};
