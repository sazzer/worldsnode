// @flow

import { request } from '../request';

test('/api/version', async () => {
    const response = await request('/api/version');
    expect(response.status).toEqual(200);
    expect(response.data).toMatchSnapshot();
});
