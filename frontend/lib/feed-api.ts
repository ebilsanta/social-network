import { GetFeedResponse } from '@/types/feed';

export class FeedAPI {
  static baseUrl =
    `${process.env.NEXT_PUBLIC_BASE_API_URL}/feeds` || 'http://localhost:8000/api/v1/feeds';

  static async getFeed(userId: string, page = 1, limit = 10): Promise<GetFeedResponse> {
    const url = new URL(`${this.baseUrl}/${userId}`);
    const params = new URLSearchParams({
      page: page.toString(),
      limit: limit.toString(),
    });
    url.search = params.toString();

    const response = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error fetching feed');
    }

    return response.json();
  }
}
