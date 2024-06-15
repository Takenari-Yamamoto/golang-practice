import { beforeEach, describe, expect, it, vi } from "vitest";
import { UseCase } from "./usecase";
import { MockUserServiceClient } from "../interface/userService";
import { MockBookServiceClient } from "../interface/bookService";

describe(UseCase, () => {
  let mockUserService: MockUserServiceClient;
  let mockBookService: MockBookServiceClient;
  let useCase: UseCase;

  beforeEach(() => {
    mockUserService = new MockUserServiceClient();
    mockBookService = new MockBookServiceClient();

    useCase = new UseCase({
      userService: mockUserService,
      bookService: mockBookService,
    });
  });

  it("should return user data with books by user id", async () => {
    vi.mocked(mockUserService.getUserById).mockResolvedValue({
      id: "user1",
      name: "John Doe",
    });
    vi.mocked(mockBookService.listBooksByUserId).mockResolvedValue([
      { id: "book1", userId: "user1", title: "Book Title 1" },
    ]);

    const result = await useCase.getUserById("user1");
    expect(result).toEqual({
      id: "user1",
      name: "John Doe",
      book: [{ id: "book1", userId: "user1", title: "Book Title 1" }],
    });

    // モックメソッドの呼び出しを検証
    expect(mockUserService.getUserById).toHaveBeenCalledWith({ id: "user1" });
    expect(mockBookService.listBooksByUserId).toHaveBeenCalledWith({
      userId: "user1",
    });
  });

  it("should return empty book list if user has no books", async () => {
    vi.mocked(mockUserService.getUserById).mockResolvedValue({
      id: "user2",
      name: "Jane Doe",
    });
    vi.mocked(mockBookService.listBooksByUserId).mockResolvedValue([]);

    const result = await useCase.getUserById("user2");

    expect(result).toEqual({
      id: "user2",
      name: "Jane Doe",
      book: [],
    });

    expect(mockUserService.getUserById).toHaveBeenCalledWith({ id: "user2" });
    expect(mockBookService.listBooksByUserId).toHaveBeenCalledWith({
      userId: "user2",
    });
  });

  it("should handle user not found", async () => {
    vi.mocked(mockUserService.getUserById).mockRejectedValue(
      new Error("User not found")
    );
    await expect(useCase.getUserById("user3")).rejects.toThrow(
      "User not found"
    );

    expect(mockUserService.getUserById).toHaveBeenCalledWith({ id: "user3" });
    expect(mockBookService.listBooksByUserId).not.toHaveBeenCalled();
  });

  it("should handle book service error", async () => {
    vi.mocked(mockUserService.getUserById).mockResolvedValue({
      id: "user4",
      name: "Alice",
    });
    vi.mocked(mockBookService.listBooksByUserId).mockRejectedValue(
      new Error("Book service error")
    );

    await expect(useCase.getUserById("user4")).rejects.toThrow(
      "Book service error"
    );

    expect(mockUserService.getUserById).toHaveBeenCalledWith({ id: "user4" });
    expect(mockBookService.listBooksByUserId).toHaveBeenCalledWith({
      userId: "user4",
    });
  });
});
