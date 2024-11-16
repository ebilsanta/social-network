import { CreateUserRequest, GetUserResponse, GetUsersResponse } from '@/types/user';

export class UserAPI {
  static baseUrl =
    `${process.env.NEXT_PUBLIC_BASE_API_URL}/users` || 'http://localhost:8000/api/v1/users';

  static async getUsers(query = '', page = 1, limit = 10): Promise<GetUsersResponse> {
    const url = new URL(this.baseUrl);
    const params = new URLSearchParams({
      query,
      page: page.toString(),
      limit: limit.toString(),
    });
    url.search = params.toString();

    const response = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error fetching users');
    }

    return response.json();
  }

  static async getUser(id: string): Promise<GetUserResponse> {
    const response = await fetch(`${this.baseUrl}/${id}`, {
      method: 'GET',
      credentials: 'include',
    });

    if (response.status === 404) {
      throw new Error('User not found');
    }

    if (!response.ok) {
      throw new Error('Error fetching user');
    }

    return response.json();
  }

  static async createUser(userData: CreateUserRequest): Promise<GetUserResponse> {
    const response = await fetch(this.baseUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error creating user');
    }

    return response.json();
  }
}
