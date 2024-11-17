import { usePagination } from '@/hooks/usePagination';
import { useUser } from '@/providers/user-provider';

export const useFeed = () => {
  const { user } = useUser();
  const { page, setMorePages, loadMoreRef } = usePagination();

  return {
    user,
    page,
    setMorePages,
    loadMoreRef,
  };
};
