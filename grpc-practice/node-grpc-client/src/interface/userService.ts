export type UserServiceClient = {
  getUserById: (request: { id: string }) => Promise<{
    id: string;
    name: string;
  }>;
};
