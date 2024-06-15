import { beforeEach, describe, expect, it, vi } from "vitest";
import { UseCase } from "./usecase";
import {
  HelloServiceClient,
  MockHelloServiceClient,
} from "../interface/helloServiceClient";

describe(UseCase, () => {
  const mockHelloServiceClient = vi.mocked(MockHelloServiceClient);
  const mock;
  const useCase = new UseCase({
    helloService: mockHelloServiceClient,
    userService: null as any,
  });
  beforeEach(() => {});

  it("should say hello", async () => {
    vi.mocked(mockHelloServiceClient.sayHello).mockResolvedValue({
      message: "Hello, Mock!",
    });

    const response = await useCase.sayHello("123");
    expect(response).toEqual({ message: "Hello, Mock!" });
    expect(mockHelloServiceClient.sayHello).toHaveBeenCalledWith({ id: "123" });
  });
});
