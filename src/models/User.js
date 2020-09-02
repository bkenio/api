const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const userSchema = new Schema(
  {
    _id: { auto: true, type: Schema.Types.ObjectId },
    password: { type: String, required: true },
    username: { type: String, required: true },
    email: {
      type: String,
      unique: true,
      required: true,
      match: /[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/,
    },
  },
  { timestamps: true }
);

userSchema.index({ username: 'text' });

module.exports = mongoose.model('User', userSchema);