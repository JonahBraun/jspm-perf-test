'use strict';

import React from 'react';
import ReactDOM from 'react-dom';

var HelloMessage = (props) => <div>Load time: {props.name}ms</div>;

ReactDOM.render(<HelloMessage name={Date.now()-performance.timing.fetchStart} />, document.getElementById('root'));
