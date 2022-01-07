'use strict';

/**
 * Sample request generator usage.
 * Contributed by jjohnsonvng:
 * https://github.com/alexfernandez/loadtest/issues/86#issuecomment-211579639
 */

const loadtest = require('loadtest');

const max_requests = 2000
const concurrency = 1000

const options = {
	url: 'http://0.0.0.0:3030',
	concurrency: concurrency,
	method: 'POST',
	body:'',
	requestsPerSecond:max_requests,
	maxSeconds:60,
	requestGenerator: (params, options, client, callback) => {
		const message = '{"title": "test","activity_group_id": "123"}';
		options.headers['Content-Length'] = message.length;
		options.headers['Content-Type'] = 'application/json';
		options.body = '{"title": "test","activity_group_id": "123"}';
		options.path = '/todo-items';
		const request = client(options, callback);
		request.write(message);
		return request;
	}
};

loadtest.loadTest(options, (error, results) => {
	if (error) {
		return console.error('Got an error: %s', error);
	}
	console.log(results);
	console.log('Tests run successfully');
});
