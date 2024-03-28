/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as z from "zod";

/**
 * A response containing information about a newly created drink.
 */
export type DrinkResponse = {
    /**
     * The name of the drink.
     */
    name?: string | undefined;
    /**
     * A brief description of the drink.
     */
    description?: string | undefined;
    /**
     * The recipe for making the drink.
     */
    recipe?: string | undefined;
    /**
     * A URL pointing to an image of the drink.
     */
    photo?: string | undefined;
};

/** @internal */
export namespace DrinkResponse$ {
    export type Inbound = {
        name?: string | undefined;
        description?: string | undefined;
        recipe?: string | undefined;
        photo?: string | undefined;
    };

    export const inboundSchema: z.ZodType<DrinkResponse, z.ZodTypeDef, Inbound> = z
        .object({
            name: z.string().optional(),
            description: z.string().optional(),
            recipe: z.string().optional(),
            photo: z.string().optional(),
        })
        .transform((v) => {
            return {
                ...(v.name === undefined ? null : { name: v.name }),
                ...(v.description === undefined ? null : { description: v.description }),
                ...(v.recipe === undefined ? null : { recipe: v.recipe }),
                ...(v.photo === undefined ? null : { photo: v.photo }),
            };
        });

    export type Outbound = {
        name?: string | undefined;
        description?: string | undefined;
        recipe?: string | undefined;
        photo?: string | undefined;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, DrinkResponse> = z
        .object({
            name: z.string().optional(),
            description: z.string().optional(),
            recipe: z.string().optional(),
            photo: z.string().optional(),
        })
        .transform((v) => {
            return {
                ...(v.name === undefined ? null : { name: v.name }),
                ...(v.description === undefined ? null : { description: v.description }),
                ...(v.recipe === undefined ? null : { recipe: v.recipe }),
                ...(v.photo === undefined ? null : { photo: v.photo }),
            };
        });
}