// @flow

/** The identity of some resource */
export type Identity = {
    id: string,
    version: string,
    created: Date,
    updated: Date
};

/** The actual description of some resource */
export type Resource<DATA> = {
    identity: Identity,
    data: DATA
};
