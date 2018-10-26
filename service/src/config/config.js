// @flow

import convict from 'convict';
import manifest from './manifest.json';

const config = convict(manifest);
config.loadFile(`./config/${config.get('env')}.json`);

config.validate({allowed: 'strict'});

export default config;
