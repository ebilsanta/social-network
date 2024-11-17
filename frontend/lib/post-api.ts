import { GetPostsByUserIdResponse } from '@/types/post';

export class PostAPI {
  static baseUrl =
    `${process.env.NEXT_PUBLIC_BASE_API_URL}/posts` || 'http://localhost:8000/api/v1/posts';

  static async getPostsByUserId(
    userId: string,
    page = 1,
    limit = 10
  ): Promise<GetPostsByUserIdResponse> {
    const url = new URL(`${this.baseUrl}/user/${userId}`);
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
      throw new Error('Error fetching posts');
    }

    return response.json();
  }
}
