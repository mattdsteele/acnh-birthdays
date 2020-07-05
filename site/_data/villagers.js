const fetch = require('node-fetch');
const cache = require('flat-cache');
const { mkdir } = require('fs');
const { resolve } = require('path');
const { promisify } = require('util');
const mkdirp = promisify(mkdir);

const getCacheKey = () => {
  const date = new Date();
  return `${date.getUTCFullYear()}-${
    date.getUTCMonth() + 1
  }-${date.getUTCDate()}`;
};
const fetchData = async () => {
  const endpoint = 'https://acnhapi.com/v1/villagers';
  const res = await fetch(endpoint);
  const json = await res.json();
  await mkdirp('./.cache');
  return json;
};
const cacheOrGet = async () => {
  const c = cache.load('villagers', resolve('./.cache'));
  const key = getCacheKey();
  const cached = c.getKey(key);
  if (cached) {
    return cached;
  }

  const villagers = await fetchData();
  c.setKey(key, villagers);
  c.save();
  return villagers;
};
module.exports = async () => {
  const unsorted = await cacheOrGet();
  return Object.keys(unsorted)
    .map((k) => unsorted[k])
    .sort((a, b) => a.name['name-USen'].localeCompare(b.name['name-USen']));
};
