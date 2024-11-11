export interface Post {
  id: string;
  caption: string;
  userId: string;
  image: string;
  createdAt: {
    seconds: number;
    nanos: number;
  };
}
