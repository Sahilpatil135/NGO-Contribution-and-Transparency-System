-- AID TYPES
INSERT INTO cause_aid_types (name, description, icon_url) VALUES
('Monetary Donations', 'Provide financial contributions to support causes, relief efforts, and organizations.', '/aid_types/monetary.png'),
('Volunteering', 'Offer your time and skills to directly assist in events, campaigns, or on-ground operations.', '/aid_types/volunteering.png'),
('Blood & Organ Donations', 'Donate blood, plasma, or organs to save lives and support critical medical needs.', '/aid_types/blood-organ.png'),
('Goods & Resources', 'Contribute essential goods like food, clothing, medicine, and other relief materials.', '/aid_types/goods-resources.png'),
('Environmental Support', 'Participate in tree planting, cleanup drives, wildlife protection, and eco-support initiatives.', '/aid_types/environmental.png'),
('Disaster Relief Assistance', 'Provide emergency aid, rescue support, shelter, and recovery assistance in disaster-hit areas.', 'https://example.com/icons/disaster-relief.png'),
('Educational Support', 'Support education through books, digital resources, mentorship, or scholarships.', 'https://example.com/icons/education.png'),
('Medical Assistance', 'Help with medicines, medical equipment, health camps, and caregiving support.', 'https://example.com/icons/medical.png');

-- DOMAINS
INSERT INTO cause_domains (name, description, icon_url) VALUES
('Urgent', 'Critical and time-sensitive causes that require immediate support and action.', '/domains/domain_example.png'),
('Strays', 'Causes focused on the rescue, care, and welfare of stray and abandoned animals.', '/domains/domain_example.png'), ('Elderly', 'Support for senior citizens including healthcare, companionship, and welfare programs.', '/domains/domain_example.png'),
('Children', 'Initiatives aimed at child welfare, safety, development, and protection.', '/domains/domain_example.png'),
('Environmental', 'Causes that focus on environmental protection, sustainability, and climate action.', '/domains/domain_example.png'),
('Specially-Abled', 'Support for differently-abled individuals through accessibility, care, and empowerment.', '/domains/domain_example.png'),
('Education', 'Programs that support learning, scholarships, literacy, and digital education access.', '/domains/domain_example.png'),
('Hunger', 'Initiatives working to eliminate hunger and provide food security for the needy.', '/domains/domain_example.png'),
('Faith', 'Faith-based aid and community support initiatives across different spiritual groups.', '/domains/domain_example.png'),
('Health', 'Causes related to medical aid, public health, disease prevention, and treatment.', '/domains/domain_example.png'),
('Poverty', 'Efforts to uplift underprivileged communities and improve living standards.', '/domains/domain_example.png'),
('Women', 'Causes that support women empowerment, safety, education, and equality.', '/domains/domain_example.png'),
('Arts & Culture', 'Initiatives promoting art, heritage preservation, creativity, and cultural development.', '/domains/domain_example.png'),
('Sports', 'Support for sports development, training, youth engagement, and athletic programs.', '/domains/domain_example.png');

-- CAUSES
INSERT INTO causes (organization_id, title, description, domain_id, aid_type_id, goal_amount, collected_amount, deadline, cover_image_url)
VALUES
-- 1 Urgent
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Immediate Flood Relief: Mumbai Slums',
  'Emergency food, shelter and medical support for families affected by recent floods.',
  (SELECT id FROM cause_domains WHERE name = 'Urgent'),
  (SELECT id FROM cause_aid_types WHERE name = 'Disaster Relief Assistance'),
  750000.00, 125000.00,
  '2025-11-30 23:59:59+05:30',
  'https://example.com/covers/urgent_flood.jpg'
),
-- 2 Strays
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'Street Vet Camp for Strays',
  'Free vaccination and sterilisation camp to control stray overpopulation and disease.',
  (SELECT id FROM cause_domains WHERE name = 'Strays'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  120000.00, 18000.00,
  '2025-09-15 23:59:59+05:30',
  'https://example.com/covers/strays_vet.jpg'
),
-- 3 Elderly
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Monthly Care Kits for Elderly',
  'Providing hygiene kits, medicines and monthly check-ups for senior citizens.',
  (SELECT id FROM cause_domains WHERE name = 'Elderly'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  200000.00, 40000.00,
  '2025-12-31 23:59:59+05:30',
  'https://example.com/covers/elderly_kits.jpg'
),
-- 4 Children
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Rural Girls Education Drive',
  'School supplies and transport support for girls in nearby villages to improve retention.',
  (SELECT id FROM cause_domains WHERE name = 'Children'),
  (SELECT id FROM cause_aid_types WHERE name = 'Educational Support'),
  300000.00, 60000.00,
  '2026-01-31 23:59:59+05:30',
  'https://example.com/covers/girls_education.jpg'
),
-- 5 Environmental
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Urban Tree Canopy Expansion',
  'Planting and maintaining trees in city pockets to reduce heat islands.',
  (SELECT id FROM cause_domains WHERE name = 'Environmental'),
  (SELECT id FROM cause_aid_types WHERE name = 'Environmental Support'),
  250000.00, 45000.00,
  '2026-03-30 23:59:59+05:30',
  'https://example.com/covers/tree_planting.jpg'
),
-- 6 Specially-Abled
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Mobility Aids for Differently-Abled',
  'Providing wheelchairs and mobility support devices for low-income recipients.',
  (SELECT id FROM cause_domains WHERE name = 'Specially-Abled'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  180000.00, 25000.00,
  '2025-10-31 23:59:59+05:30',
  'https://example.com/covers/mobility_aids.jpg'
),
-- 7 Education
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'After-School Tutoring Program',
  'Volunteer-driven tuition for children in slum communities to improve learning outcomes.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  90000.00, 12000.00,
  '2026-02-28 23:59:59+05:30',
  'https://example.com/covers/tutoring.jpg'
),
-- 8 Hunger
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Community Meal Kits Distribution',
  'Monthly food kit distribution to vulnerable households in the city.',
  (SELECT id FROM cause_domains WHERE name = 'Hunger'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  150000.00, 30000.00,
  '2025-12-15 23:59:59+05:30',
  'https://example.com/covers/meal_kits.jpg'
),
-- 9 Faith
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Interfaith Community Kitchen',
  'Support the community kitchen serving free meals irrespective of faith.',
  (SELECT id FROM cause_domains WHERE name = 'Faith'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  60000.00, 5000.00,
  '2025-11-15 23:59:59+05:30',
  'https://example.com/covers/community_kitchen.jpg'
),
-- 10 Health
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Rural Health Camp: Mobile Clinic',
  'Weekend mobile clinic offering screenings, medicines and referrals in remote villages.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  220000.00, 40000.00,
  '2025-10-20 23:59:59+05:30',
  'https://example.com/covers/health_camp.jpg'
),
-- 11 Poverty
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Small Business Grants for Women',
  'Seed grants to help women start micro-enterprises and lift families out of poverty.',
  (SELECT id FROM cause_domains WHERE name = 'Poverty'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  500000.00, 70000.00,
  '2026-04-30 23:59:59+05:30',
  'https://example.com/covers/micro_grants.jpg'
),
-- 12 Women
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Safe Space & Legal Aid for Women',
  'Set up a helpline and legal clinic for gender-based violence survivors.',
  (SELECT id FROM cause_domains WHERE name = 'Women'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  275000.00, 35000.00,
  '2025-12-31 23:59:59+05:30',
  'https://example.com/covers/legal_aid.jpg'
),
-- 13 Arts & Culture
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Village Arts Revival Program',
  'Support traditional craftspeople with training and micro-sales channels.',
  (SELECT id FROM cause_domains WHERE name = 'Arts & Culture'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  110000.00, 15000.00,
  '2026-05-31 23:59:59+05:30',
  'https://example.com/covers/arts_culture.jpg'
),
-- 14 Sports
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Girls Sports Coaching & Kits',
  'Provide coaching, equipment and tournament fees to encourage girls in sports.',
  (SELECT id FROM cause_domains WHERE name = 'Sports'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  140000.00, 20000.00,
  '2026-03-31 23:59:59+05:30',
  'https://example.com/covers/sports_kits.jpg'
),
-- 15 Additional: Blood & Organ Donations (ensures this aid_type is present)
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'Blood Drive: Citywide Donation Camp',
  'Partnering with hospitals to create a donor registry and host drives.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Blood & Organ Donations'),
  50000.00, 8000.00,
  '2025-11-10 23:59:59+05:30',
  'https://example.com/covers/blood_drive.jpg'
),
-- 16 Additional: Educational Support (ensures this aid_type is present)
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Digital Classroom: Tablets for Rural Kids',
  'Provide tablets and offline learning content for children with no internet access.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Educational Support'),
  320000.00, 45000.00,
  '2026-06-30 23:59:59+05:30',
  'https://example.com/covers/tablets.jpg'
),
-- 17 Additional: Volunteering example (ensure multiple volunteering entries)
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Beach Cleanup & Awareness Drive',
  'Volunteer-driven beach cleanup with local schools and businesses.',
  (SELECT id FROM cause_domains WHERE name = 'Environmental'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  40000.00, 5000.00,
  '2025-12-01 23:59:59+05:30',
  'https://example.com/covers/beach_cleanup.jpg'
),
-- 18 Additional: Goods & Resources (another example)
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Winter Blankets for Street Families',
  'Distribution of warm blankets and basic winter clothing to those in need.',
  (SELECT id FROM cause_domains WHERE name = 'Poverty'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  90000.00, 15000.00,
  '2025-11-30 23:59:59+05:30',
  'https://example.com/covers/blankets.jpg'
),
-- 19 Additional: Environmental Support (another example)
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'School Composting Project',
  'Install compost pits and run workshops to reduce organic waste in schools.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Environmental Support'),
  60000.00, 8000.00,
  '2026-02-28 23:59:59+05:30',
  'https://example.com/covers/compost.jpg'
),
-- 20 Additional: Disaster Relief Assistance (another example)
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Drought Relief: Water Tanker & Bore Repairs',
  'Immediate water supply and borewell repairs for drought-affected villages.',
  (SELECT id FROM cause_domains WHERE name = 'Urgent'),
  (SELECT id FROM cause_aid_types WHERE name = 'Disaster Relief Assistance'),
  300000.00, 60000.00,
  '2025-10-31 23:59:59+05:30',
  'https://example.com/covers/drought_relief.jpg'
),
-- 21 Additional: Monitory Donations (another example)
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Scholarships for Women in Tech',
  'Funds to support meritorious girls pursuing technical degrees.',
  (SELECT id FROM cause_domains WHERE name = 'Women'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  350000.00, 90000.00,
  '2026-07-31 23:59:59+05:30',
  'https://example.com/covers/scholarships.jpg'
),
-- 22 Additional: Medical Assistance (another example)
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Clean Water & Sanitation Health Drive',
  'Combines water purification distribution with hygiene education and basic treatment.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  260000.00, 32000.00,
  '2026-01-31 23:59:59+05:30',
  'https://example.com/covers/wash_health.jpg'
),
-- 23 Additional: Goods & Resources for Children
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'School Stationery Packs for Underprivileged Kids',
  'Provide stationery, bags and summer uniforms for local municipal school children.',
  (SELECT id FROM cause_domains WHERE name = 'Children'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  50000.00, 7000.00,
  '2025-10-15 23:59:59+05:30',
  'https://example.com/covers/stationery.jpg'
),
-- 24 Additional: Arts & Culture volunteering
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Heritage Wall Mural Project',
  'Volunteer artists working with local youth to create heritage murals.',
  (SELECT id FROM cause_domains WHERE name = 'Arts & Culture'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  45000.00, 6000.00,
  '2026-04-30 23:59:59+05:30',
  'https://example.com/covers/mural.jpg'
);
