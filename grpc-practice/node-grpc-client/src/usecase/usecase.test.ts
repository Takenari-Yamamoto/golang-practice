import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { UseCase } from "./usecase";
import {
  HelloServiceClient,
  MockServiceClient,
} from "../interface/helloServiceClient";

describe(UseCase, () => {
  let usecase: UseCase;
  let mockServiceClient: MockServiceClient;

  beforeEach(() => {
    mockServiceClient = new MockServiceClient();
    usecase = new UseCase({ helloService: mockServiceClient });
  });

  it("should say hello", async () => {
    vi.spyOn(mockServiceClient, "sayHello").mockResolvedValueOnce({
      message: "Hello",
    });

    await usecase.sayHello("123");

    expect(mockServiceClient.sayHello).toHaveBeenCalledWith({ id: "123" });
    expect(mockServiceClient.sayHello).toHaveBeenCalledTimes(1);
  });
});
