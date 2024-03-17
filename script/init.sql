INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('deposit_1', '{ \"name\": \"10 Coins\", \"price\": \"0.99\", \"coin\": 10 }', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('deposit_2', '{ \"name\": \"30 Coins\", \"price\": \"2.97\", \"coin\": 30 }', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('deposit_3', '{ \"name\": \"50 Coins\", \"price\": \"4.90\", \"coin\": 50 }', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('deposit_4', '{ \"name\": \"100 Coins\", \"price\": \"9.88\", \"coin\": 100 }', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('payway_1', '{ \"name\": \"paypal\"}', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('payway_2', '{ \"name\": \"credit\"}', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('tips_1', 'Payment Notice', 1);


INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_1', 'outfit', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_2', 'backless outfit', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_3', 'underboob cutout', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_4', 'shoulder cutout', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_5', 'sweater', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_6', 'off-shoulder sweater', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_7', 'halter dress', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_8', 'tank top', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_9', 'strapless leotard', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_10', 'leotard', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_11', 'sportswear', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_12', 'swimsuit', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_13', 'bikini', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_14', 'business suit', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_15', 'gothic_lolita', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_16', 'Maid dress', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_17', 'white_windbreaker', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_18', 'overcoat', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_19', 'coat', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_21', 'sweet_lolita', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_22', 'zentai', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_23', 'Raincoat', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_24', 'pajamas', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_25', 'sweatshirt', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_26', 'hoodie', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_27', 'Sailor Uniform', 1);
INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('cloth_28', 'Soccer Jersey', 1);

INSERT INTO `qnc_kv` (`name`, `value`, `status`) VALUES ('regions', '[\"ID\", \"US\"]', 1);


INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('US', 10.00, 'USD', '$', 1.00, 0.99, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('US', 30.00, 'USD', '$', 3.00, 2.97, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('US', 50.00, 'USD', '$', 5.00, 4.95, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('US', 100.00, 'USD', '$', 10.00,9.88, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('ID', 10.00, 'IDR', 'Rp', 15600, 15555, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('ID', 30.00, 'IDR', 'Rp', 46860, 46666, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('ID', 50.00, 'IDR', 'Rp', 78100, 77777, NULL, 1708785940, 1708785968);
INSERT INTO `qnc_deposit_config` (`country_code`, `coins`, `currency`, `sign`, `price`, `actual_price`, `remark`, `create_time`, `update_time`) VALUES ('ID', 100.00, 'IDR', 'Rp', 157000,155555, NULL, 1708785940, 1708785968);


INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('USD', 0.13897755, 0.13897755, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('IDR', 2174.75693428, 2174.75693428, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('SGD', 0.18593748, 0.18593748, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('THB', 4.97878414, 4.97878414, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('PHP', 7.72942837, 7.72942837, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('VND', 3434.91671119, 3434.91671119, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('MYR', 0.65381949, 0.65381949, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('LAK', 2902.03702569, 2902.03702569, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('KHR', 563.02549074, 563.02549074, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('MMK', 291.82963928, 291.82963928, '2024-03-16', '', 1710610214, 1710610328);
INSERT INTO `qnc_currency_record` (`currency`, `base_xt`, `latest_xt`, `update_date`, `remark`, `create_time`, `update_time`) VALUES ('INR', 11.52048758, 11.52048758, '2024-03-16', '', 1710610214, 1710610328);