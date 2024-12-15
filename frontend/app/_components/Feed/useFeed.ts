import { usePathname } from 'next/navigation';
import { usePagination } from '@/hooks/usePagination';
import { useUser } from '@/providers/user-provider';

export const useFeed = () => {
  const { user } = useUser();
  const { page, setMorePages, loadMoreRef } = usePagination();
  const path = usePathname();
  const postId = path.split('p/')[1];

  return {
    user,
    page,
    setMorePages,
    loadMoreRef,
    postId,
  };
};
