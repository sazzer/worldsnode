// @flow

import { request } from '../request';

test('/health', async () => {
    const response = await request('/health');
    expect(response.status).toEqual(200);
    expect(response.data.status).toEqual("pass");
    // expect(response.data.details['system:uptime'][0].status).toEqual('pass');
    // expect(response.data.details['database:ping'][0].status).toEqual('pass');
});
