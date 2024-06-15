import { Mock, vi } from "vitest";

type Primitive = string | number | boolean | undefined | null;

export type ObjectWithMocks<Type> = {
  [Property in keyof Type]: Type[Property] extends CallableFunction
    ? Mock
    : Type[Property] extends Record<string, unknown>
    ? ObjectWithMocks<Type[Property]>
    : Type[Property];
};

const primitiveTypes = new Set(["string", "boolean", "number", "undefined"]);

const mockApi = (): Mock => vi.fn();

export const deepCopyWithMocks = <T>(original: T): ObjectWithMocks<T> => {
  const copy = {} as Record<string, unknown>;

  for (const key in original) {
    const value = original[key];
    if (typeof value === "function") {
      copy[key] = vi.fn();
    } else if (primitiveTypes.has(typeof value)) {
      copy[key] = value;
    } else {
      copy[key] = deepCopyWithMocks(value);
    }
  }

  return copy as ObjectWithMocks<T>;
};

class MockClient<T> {
  public mock: ObjectWithMocks<T>;

  constructor() {
    this.mock = deepCopyWithMocks({} as T);
  }

  getMock() {
    return this.mock;
  }
}

export { MockClient };
