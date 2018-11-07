// @flow

/** Type representing a Problem, as defined in RFC-7807 */
export type ProblemType = {
    type: string,
    title: string,
    status: number,
    detail?: string,
    instance?: string,
};
