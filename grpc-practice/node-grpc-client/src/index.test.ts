import { add } from "../src/index";
import { expect, describe, it } from "vitest";

describe("add function", () => {
  it("should return the sum of two numbers", () => {
    expect(add(1, 2)).toBe(1);
  });
});
