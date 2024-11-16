import useSWR from 'swr';
import { UserAPI } from '@/lib/user-api';
import { useUser } from '@/providers/user-provider';

export const useProfileCard = () => {
  const { user } = useUser();
  const fetchUser = async (userId: string) => {
    const response = await UserAPI.getUser(userId);
    return response.data;
  };
  const { data, error } = useSWR(
    user?.id ? [`/api/users/${user.id}`, user.id] : null,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    ([_, userId]) => fetchUser(userId)
  );
  return {
    data,
    error,
  };
};
