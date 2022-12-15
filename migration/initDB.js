use users
db.users.stats()
db.users.createIndex({ email: 1 });
db.users.createIndex({ '$**': 'text' });
db.users.getIndexes();

use groups
db.groups.stats()
db.groups.createIndex({ creator_id: 1 });
db.groups.createIndex({ '$**': 'text' });
db.groups.getIndexes();

use group_memberships
db.group_memberships.stats()
db.group_memberships.createIndex({ group_id: 1, user_id: 1 });
db.group_memberships.createIndex({ '$**': 'text' });
db.group_memberships.getIndexes();