export type BookServiceClient = {
  listBooksByUserId: (request: { userId: string }) => Promise<{
    id: string;
    userId: string;
    title: string;
  }>;
};
