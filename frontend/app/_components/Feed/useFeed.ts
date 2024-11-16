import { useEffect, useRef, useState } from 'react';
import { useUser } from '@/providers/user-provider';

export const useFeed = () => {
  const { user } = useUser();
  const [page, setPage] = useState(1);
  const [morePages, setMorePages] = useState(true);
  const loadMoreRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (morePages && entry.isIntersecting) {
          setPage((prevPage) => prevPage + 1);
        }
      },
      {
        rootMargin: '200px',
        threshold: 1.0,
      }
    );

    if (loadMoreRef.current) {
      observer.observe(loadMoreRef.current);
    }

    return () => {
      if (loadMoreRef.current) {
        observer.unobserve(loadMoreRef.current);
      }
    };
  }, [morePages]);

  return {
    user,
    page,
    setMorePages,
    loadMoreRef,
  };
};
