/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AccountHandle } from "./accountHandle";
import type { AccountName } from "./accountName";
import type { TagList } from "./tagList";

export type PublicProfileAllOf = {
  handle?: AccountHandle;
  name?: AccountName;
  bio?: string;
  image?: string;
  createdAt?: string;
  interests?: TagList;
};
