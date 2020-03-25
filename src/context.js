const jwt = require('jsonwebtoken');

const users = require('./loaders/users');
const videos = require('./loaders/videos');
const uploads = require('./loaders/uploads');

module.exports = ({ req }) => {
  let user = null;
  const token = req.headers.authorization || '';
  if (token) {
    if (token.includes('SERVICE')) {
      if (!token.split('SERVICE ')[1] === process.env.CONVERSION_API_KEY) {
        throw new Error('invalid service api key');
      }
    } else {
      user = jwt.verify(token.split(' ')[1], process.env.JWT_KEY);
    }
  }

  return {
    user,
    users,
    videos,
    uploads,
  };
};