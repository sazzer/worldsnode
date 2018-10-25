// @flow

import { buildService } from './service';
import { config } from './config';

const port = config.get('http.port');

const app = buildService();
app.listen(port, () => console.log(`Example app listening on port ${port}!`));
