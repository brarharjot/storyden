/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { Identifier } from "./identifier";

export interface Category {
  id?: Identifier;
  name?: string;
  description?: string;
  colour?: string;
  sort?: number;
  admin?: boolean;
  postCount?: number;
}
