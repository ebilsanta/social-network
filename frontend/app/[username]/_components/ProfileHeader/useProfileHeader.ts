import useSWR from 'swr';
import { UserAPI } from '@/lib/user-api';

export const useProfileHeader = (username?: string) => {
  const fetchUser = async () => {
    const response = await UserAPI.getUserByUsername(username!);
    return response.data;
  };
  const { data, error } = useSWR(username ? `/api/users/username/${username}` : null, fetchUser);
  return {
    data,
    error,
  };
};
