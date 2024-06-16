import { vi } from "vitest";

export type BookServiceClient = {
  listBooksByUserId: (request: { userId: string }) => Promise<
    {
      id: string;
      userId: string;
      title: string;
    }[]
  >;

  findBookById: (request: { id: string }) => Promise<{
    id: string;
    userId: string;
    title: string;
  }>;

  createBook: (request: { userId: string; title: string }) => Promise<{
    id: string;
    userId: string;
    title: string;
  }>;

  updateBook: (request: { id: string; title: string }) => Promise<{
    id: string;
    userId: string;
    title: string;
  }>;
};

export class MockBookServiceClient implements BookServiceClient {
  listBooksByUserId = vi.fn<
    [request: { userId: string }],
    Promise<{ id: string; userId: string; title: string }[]>
  >();

  findBookById = vi.fn<
    [request: { id: string }],
    Promise<{ id: string; userId: string; title: string }>
  >();

  createBook = vi.fn<
    [request: { userId: string; title: string }],
    Promise<{ id: string; userId: string; title: string }>
  >();
  updateBook = vi.fn<
    [request: { id: string; title: string }],
    Promise<{ id: string; userId: string; title: string }>
  >();
}
