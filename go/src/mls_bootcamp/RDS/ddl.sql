CREATE TABLE `pm2_daily` (
  `dt` varchar(8) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL, -- 일자
  `region` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL, -- 위치
  `no2` float DEFAULT NULL, -- 이산화질소농도
  `o3` float DEFAULT NULL, -- 오존농도
  `co` float DEFAULT NULL, -- 일산화탄소농도
  `so2` float DEFAULT NULL, -- 아황산가스
  `pm10` float DEFAULT NULL, -- 미세먼지
  `spm25` float DEFAULT NULL, -- 초미세먼지

  PRIMARY KEY (`dt`,`region`)
) DEFAULT CHARSET=utf8
;

