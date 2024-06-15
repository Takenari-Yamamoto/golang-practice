import { vi } from "vitest";

export type UserServiceClient = {
  getUserById: (request: { id: string }) => Promise<{
    id: string;
    name: string;
  }>;
  listUsers: () => Promise<
    {
      id: string;
      name: string;
    }[]
  >;
  createUser: (request: { name: string }) => Promise<{
    id: string;
    name: string;
  }>;
  updateUser: (request: { id: string; name: string }) => Promise<{
    id: string;
    name: string;
  }>;
  deleteUser: (request: { id: string }) => Promise<{
    id: string;
    name: string;
  }>;
};

export class MockUserServiceClient implements UserServiceClient {
  getUserById = vi.fn<
    [request: { id: string }],
    Promise<{ id: string; name: string }>
  >() as (request: { id: string }) => Promise<{ id: string; name: string }>;
  listUsers = vi.fn<
    [],
    Promise<{ id: string; name: string }[]>
  >() as () => Promise<{ id: string; name: string }[]>;
  createUser = vi.fn<
    [request: { name: string }],
    Promise<{ id: string; name: string }>
  >() as (request: { name: string }) => Promise<{ id: string; name: string }>;
  updateUser = vi.fn<
    [request: { id: string; name: string }],
    Promise<{ id: string; name: string }>
  >() as (request: {
    id: string;
    name: string;
  }) => Promise<{ id: string; name: string }>;
  deleteUser = vi.fn<
    [request: { id: string }],
    Promise<{ id: string; name: string }>
  >() as (request: { id: string }) => Promise<{ id: string; name: string }>;
}
