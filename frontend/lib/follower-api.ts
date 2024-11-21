export class FollowerAPI {
  static baseUrl =
    `${process.env.NEXT_PUBLIC_BASE_API_URL}/followers` || 'http://localhost:8000/api/v1/followers';

  static async addFollower(followerId: string, followingId: string): Promise<void> {
    const response = await fetch(this.baseUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        followerId,
        followingId,
      }),
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error following user');
    }

    return response.json();
  }

  static async deleteFollower(followerId: string, followingId: string): Promise<void> {
    const response = await fetch(this.baseUrl, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        followerId,
        followingId,
      }),
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error unfollowing user');
    }

    return response.json();
  }
}
