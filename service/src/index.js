// @flow

import config from './config';
import buildService from './service';
import versionHandler from './service/version';
import uptimeHealthchecks from './health/system/uptimeHealthcheck';
import buildHealthchecksHandler from './health/rest';

const service = buildService([
    versionHandler,
    buildHealthchecksHandler([uptimeHealthchecks]),
]);
const port = config.get('http.port');
/* eslint-disable */
service.listen(port, () => console.log(`Service listening on port ${port}`));
