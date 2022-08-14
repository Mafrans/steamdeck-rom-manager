import { ComponentChildren } from "preact";

export type Props<T> = {
  className?: string;
  children?: ComponentChildren;
} & T;
