import { vi } from "vitest";

export type HelloServiceClient = {
  readonly sayHello: (request: { id: string }) => Promise<{
    message: string;
  }>;
};

export class MockServiceClient implements HelloServiceClient {
  sayHello = vi.fn();
}
