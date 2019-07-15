/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2019/7/2 10:57:32                            */
/*==============================================================*/


drop table if exists Good;

drop table if exists Relationship_5;

drop table if exists chat_log;

drop table if exists comment;

drop table if exists comment_picture;

drop table if exists complaint;

drop table if exists evaluate;

drop table if exists good_picture;

drop table if exists `release`;

drop table if exists tag;

drop table if exists transaction;

drop table if exists user;

/*==============================================================*/
/* Table: Good                                                  */
/*==============================================================*/
create table Good
(
   release_id           int not null,
   good_name            varchar(64) not null,
   good_description     varchar(1024) not null,
   primary key (release_id)
);

/*==============================================================*/
/* Table: Relationship_5                                        */
/*==============================================================*/
create table Relationship_5
(
   release_id           int not null,
   tag_id               int not null,
   primary key (release_id, tag_id)
);

/*==============================================================*/
/* Table: chat_log                                              */
/*==============================================================*/
create table chat_log
(
   chat_log_id          int not null,
   user_id              int not null,
   use_user_id          int not null,
   mongo_chat_id        varchar(1024) not null,
   primary key (chat_log_id)
);

/*==============================================================*/
/* Table: comment                                               */
/*==============================================================*/
create table comment
(
   comment_id           varchar(1024) not null,
   user_id              int not null,
   release_id           int not null,
   content              varchar(1024) not null,
   mongo_comment_id     varchar(1024),
   primary key (comment_id)
);

/*==============================================================*/
/* Table: comment_picture                                       */
/*==============================================================*/
create table comment_picture
(
   comment_picture_id   int not null,
   comment_id           varchar(1024) not null,
   comment_picture_url  varchar(1024) not null,
   primary key (comment_picture_id)
);

/*==============================================================*/
/* Table: complaint                                             */
/*==============================================================*/
create table complaint
(
   complaint_id         int not null,
   user_id              int,
   use_user_id          int,
   content              varchar(1024) not null,
   primary key (complaint_id)
);

/*==============================================================*/
/* Table: evaluate                                              */
/*==============================================================*/
create table evaluate
(
   evaluate_id          int not null,
   user_id              int not null,
   score                int not null,
   primary key (evaluate_id)
);

/*==============================================================*/
/* Table: good_picture                                          */
/*==============================================================*/
create table good_picture
(
   good_picture_id      int not null,
   release_id           int,
   good_picture_url     varchar(1024) not null,
   primary key (good_picture_id)
);

/*==============================================================*/
/* Table: "release"                                             */
/*==============================================================*/
create table `release`
(
   release_id           int not null,
   user_id              int not null,
   Goo_release_id       int,
   tra_release_id       int,
   is_saled             bool not null,
   release_time         datetime not null,
   valid_time           datetime not null,
   is_appointed         bool not null,
   is_finished          bool not null,
   primary key (release_id)
);

/*==============================================================*/
/* Table: tag                                                   */
/*==============================================================*/
create table tag
(
   tag_id               int not null,
   tag_name             varchar(16) not null,
   primary key (tag_id)
);

/*==============================================================*/
/* Table: transaction                                           */
/*==============================================================*/
create table transaction
(
   release_id           int not null,
   user_id              int not null,
   create_time          datetime not null,
   primary key (release_id)
);

/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
   user_id              int not null,
   user_name            varchar(16) not null,
   password             varchar(16) not null,
   portrait_url         varchar(1024) not null,
   telephone            char(11) not null,
   primary key (user_id)
);

alter table Good add constraint FK_Relationship_6 foreign key (release_id)
      references `release` (release_id) on delete restrict on update restrict;

alter table Relationship_5 add constraint FK_Relationship_7 foreign key (release_id)
      references Good (release_id) on delete restrict on update restrict;

alter table Relationship_5 add constraint FK_Relationship_8 foreign key (tag_id)
      references tag (tag_id) on delete restrict on update restrict;

alter table chat_log add constraint FK_Relationship_14 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table chat_log add constraint FK_Relationship_15 foreign key (use_user_id)
      references user (user_id) on delete restrict on update restrict;

alter table comment add constraint FK_Relationship_12 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table comment add constraint FK_Relationship_13 foreign key (release_id)
      references transaction (release_id) on delete restrict on update restrict;

alter table comment_picture add constraint FK_Relationship_11 foreign key (comment_id)
      references comment (comment_id) on delete restrict on update restrict;

alter table complaint add constraint FK_Relationship_17 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table complaint add constraint FK_Relationship_18 foreign key (use_user_id)
      references user (user_id) on delete restrict on update restrict;

alter table evaluate add constraint FK_Relationship_16 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table good_picture add constraint FK_Relationship_10 foreign key (release_id)
      references Good (release_id) on delete restrict on update restrict;

alter table `release` add constraint FK_Relationship_4 foreign key (tra_release_id)
      references transaction (release_id) on delete restrict on update restrict;

alter table `release` add constraint FK_Relationship_5 foreign key (Goo_release_id)
      references Good (release_id) on delete restrict on update restrict;

alter table `release` add constraint FK_Relationship_9 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table transaction add constraint FK_Relationship_1 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table transaction add constraint FK_Relationship_3 foreign key (release_id)
      references `release` (release_id) on delete restrict on update restrict;