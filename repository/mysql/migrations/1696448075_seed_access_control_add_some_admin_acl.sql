-- +migrate Up
INSERT INTO `permissions`(`id`, `title`) values(1,'user-list');
INSERT INTO `permissions`(`id`, `title`) values(2,'user-delete');

INSERT INTO `access_controls`(`actor_type`, `actor_id`, `permission_id`) values('role',2,1);
INSERT INTO `access_controls`(`actor_type`, `actor_id`, `permission_id`) values('role',2,2);

-- +migrate Down
DELETE FROM `permissions` WHERE `id` IN (1,2);
