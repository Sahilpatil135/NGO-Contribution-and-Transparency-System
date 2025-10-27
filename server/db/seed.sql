-- AID TYPES
INSERT INTO cause_aid_types (name, description, icon_url) VALUES
('Monetary Donations', 'Provide financial contributions to support causes, relief efforts, and organizations.', '/aid_types/monetary_donations.png'),
('Volunteering', 'Offer your time and skills to directly assist in events, campaigns, or on-ground operations.', '/aid_types/volunteering.png'),
('Blood & Organ Donations', 'Donate blood, plasma, or organs to save lives and support critical medical needs.', '/aid_types/blood_organ_donations.png'),
('Goods & Resources', 'Contribute essential goods like food, clothing, medicine, and other relief materials.', '/aid_types/goods_resources.png'),
('Environmental Support', 'Participate in tree planting, cleanup drives, wildlife protection, and eco-support initiatives.', '/aid_types/environmental_support.png'),
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
  'Heavy monsoon flooding has submerged densely populated informal settlements in Mumbai, leaving families without access to shelter, food, or medical care. This initiative aims to deliver emergency food kits, medical aid, and temporary housing support. Rapid intervention is critical to prevent disease outbreaks and further displacement.',
  (SELECT id FROM cause_domains WHERE name = 'Urgent'),
  (SELECT id FROM cause_aid_types WHERE name = 'Disaster Relief Assistance'),
  750000.00, 125000.00,
  '2026-12-31 23:59:59+05:30',
  'https://example.com/covers/urgent_flood.jpg'
),
-- 2 Strays
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'Street Vet Camp for Strays',
  'Unvaccinated and unsterilised stray animals continue to multiply, increasing the risk of rabies and other infections in urban neighborhoods. This camp will provide on-ground vaccination, sterilisation, and basic treatment for injured animals. It also aims to spread animal welfare awareness among local residents.',
  (SELECT id FROM cause_domains WHERE name = 'Strays'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  120000.00, 18000.00,
  '2026-11-30 23:59:59+05:30',
  'https://example.com/covers/strays_vet.jpg'
),
-- 3 Elderly
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Monthly Care Kits for Elderly',
  'Many senior citizens living alone or in low-income households struggle to afford medicines and hygiene essentials. This initiative provides monthly doorstep care kits containing nutritional support, personal hygiene items, and basic medication. It also includes periodic health check-ups by volunteer medical professionals.',
  (SELECT id FROM cause_domains WHERE name = 'Elderly'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  200000.00, 40000.00,
  '2026-12-31 23:59:59+05:30',
  'https://example.com/covers/elderly_kits.jpg'
),
-- 4 Children
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Rural Girls Education Drive',
  'A large number of rural girls discontinue school due to lack of transport, uniforms, and basic educational support. This program funds school supplies, safe transportation, and mentorship to help girls continue their education with dignity. It aims to significantly improve female retention in secondary education.',
  (SELECT id FROM cause_domains WHERE name = 'Children'),
  (SELECT id FROM cause_aid_types WHERE name = 'Educational Support'),
  300000.00, 60000.00,
  '2027-01-31 23:59:59+05:30',
  'https://example.com/covers/girls_education.jpg'
),
-- 5 Environmental
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Urban Tree Canopy Expansion',
  'Indian cities are facing rising urban heat and air pollution due to declining green cover. This project focuses on planting native, climate-resilient trees in heat-prone zones and maintaining them until maturity. Trained volunteers will ensure long-term care to guarantee survival and impact.',
  (SELECT id FROM cause_domains WHERE name = 'Environmental'),
  (SELECT id FROM cause_aid_types WHERE name = 'Environmental Support'),
  250000.00, 45000.00,
  '2027-03-30 23:59:59+05:30',
  'https://example.com/covers/tree_planting.jpg'
),
-- 6 Specially-Abled
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Mobility Aids for Differently-Abled',
  'Many low-income differently-abled individuals remain immobile due to lack of access to wheelchairs and walking aids. This initiative aims to distribute durable mobility devices, customised as per medical needs, for improved independence and dignity. It prioritises those from underprivileged rural and urban areas.',
  (SELECT id FROM cause_domains WHERE name = 'Specially-Abled'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  180000.00, 25000.00,
  '2026-11-30 23:59:59+05:30',
  'https://example.com/covers/mobility_aids.jpg'
),
-- 7 Education
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'After-School Tutoring Program',
  'Children from underserved urban communities often lack academic support after school. This volunteer-led program offers daily classes focusing on foundational subjects like math, language, and science. It aims to improve learning outcomes and reduce dropout rates among first-generation learners.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  90000.00, 12000.00,
  '2027-02-28 23:59:59+05:30',
  'https://example.com/covers/tutoring.jpg'
),
-- 8 Hunger
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Community Meal Kits Distribution',
  'Many urban households face food insecurity due to unstable income and rising costs. This program ensures monthly distribution of nutrient-balanced food kits containing essentials like rice, lentils, and cooking oil. It prioritises widows, elderly persons, and daily-wage families in high-risk zones.',
  (SELECT id FROM cause_domains WHERE name = 'Hunger'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  150000.00, 30000.00,
  '2026-12-15 23:59:59+05:30',
  'https://example.com/covers/meal_kits.jpg'
),
-- 9 Faith
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Interfaith Community Kitchen',
  'An inclusive community kitchen that serves free meals daily to homeless individuals regardless of religion, caste, or background. This initiative fosters harmony while ensuring basic nutrition for all. Donations help sustain daily operations and expand to more localities.',
  (SELECT id FROM cause_domains WHERE name = 'Faith'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  60000.00, 5000.00,
  '2026-11-15 23:59:59+05:30',
  'https://example.com/covers/community_kitchen.jpg'
),
-- 10 Health
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Rural Health Camp: Mobile Clinic',
  'Remote villages often lack access to primary healthcare and essential medicines. This mobile clinic will conduct weekend medical check-ups, distribute medicines, and provide referrals for critical cases. The initiative is designed to serve economically vulnerable households in hard-to-reach regions.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  220000.00, 40000.00,
  '2026-12-20 23:59:59+05:30',
  'https://example.com/covers/health_camp.jpg'
),
-- 11 Poverty
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Small Business Grants for Women',
  'Thousands of low-income women possess strong entrepreneurial potential but lack seed funding to begin small businesses. This project offers micro-grants, financial literacy guidance, and mentorship to empower them toward long-term financial independence. Preference is given to single mothers and households below the poverty line.',
  (SELECT id FROM cause_domains WHERE name = 'Poverty'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  500000.00, 70000.00,
  '2027-04-30 23:59:59+05:30',
  'https://example.com/covers/micro_grants.jpg'
),
-- 12 Women
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Safe Space & Legal Aid for Women',
  'Survivors of gender-based violence often lack legal protection and mental health support. This initiative establishes a helpline, legal consultation center, and safe transitional space for affected women. The program also connects them with counsellors and long-term rehabilitation assistance.',
  (SELECT id FROM cause_domains WHERE name = 'Women'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  275000.00, 35000.00,
  '2026-12-31 23:59:59+05:30',
  'https://example.com/covers/legal_aid.jpg'
),
-- 13 Arts & Culture
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Village Arts Revival Program',
  'Traditional artisans in rural India face declining demand, threatening cultural heritage and livelihoods. This initiative supports training, digital cataloging, and sustainable sales channels to revive heritage crafts. It also aims to connect artisans with fair trade networks for global reach.',
  (SELECT id FROM cause_domains WHERE name = 'Arts & Culture'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  110000.00, 15000.00,
  '2027-05-31 23:59:59+05:30',
  'https://example.com/covers/arts_culture.jpg'
),
-- 14 Sports
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Girls Sports Coaching & Kits',
  'Young girls in rural and semi-urban areas rarely receive structured sports training or professional coaching. This program provides quality sports equipment, certified trainers, and funds tournament participation to build confidence and leadership. It promotes fitness, teamwork, and future career opportunities.',
  (SELECT id FROM cause_domains WHERE name = 'Sports'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  140000.00, 20000.00,
  '2027-03-31 23:59:59+05:30',
  'https://example.com/covers/sports_kits.jpg'
),
-- 15 Blood & Organ Donations
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'Blood Drive: Citywide Donation Camp',
  'Increasing emergency medical cases have led to severe blood shortages in metropolitan hospitals. This initiative sets up blood donation camps and creates a verified donor registry for quicker response during emergencies. It collaborates with certified hospitals to ensure safe and transparent medical handling.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Blood & Organ Donations'),
  50000.00, 8000.00,
  '2026-11-10 23:59:59+05:30',
  'https://example.com/covers/blood_drive.jpg'
),
-- 16 Educational Support
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Digital Classroom: Tablets for Rural Kids',
  'Children in remote villages lack access to digital learning despite curriculum modernization. This project distributes tablets pre-loaded with high-quality offline educational content to support self-paced learning. It also trains teachers and parents on responsible device usage for long-term impact.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Educational Support'),
  320000.00, 45000.00,
  '2027-06-30 23:59:59+05:30',
  'https://example.com/covers/tablets.jpg'
),
-- 17 Volunteering
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Beach Cleanup & Awareness Drive',
  'Plastic waste on coastlines poses a severe threat to marine life and local tourism. This drive mobilizes volunteers to clean beaches and conduct awareness workshops for schools and citizens. The project also collaborates with municipal bodies for responsible waste management.',
  (SELECT id FROM cause_domains WHERE name = 'Environmental'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  40000.00, 5000.00,
  '2026-12-01 23:59:59+05:30',
  'https://example.com/covers/beach_cleanup.jpg'
),
-- 18 Goods & Resources
(
  (SELECT id FROM organizations WHERE organization_name = 'Seva Sagar Trust'),
  'Winter Blankets for Street Families',
  'Homeless communities and daily-wage workers are at high risk during extreme winter nights in northern India. This campaign distributes insulated blankets and warm clothing to vulnerable families living in open or temporary shelters. Priority is given to elderly individuals and children.',
  (SELECT id FROM cause_domains WHERE name = 'Poverty'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  90000.00, 15000.00,
  '2026-11-30 23:59:59+05:30',
  'https://example.com/covers/blankets.jpg'
),
-- 19 Environmental Support
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'School Composting Project',
  'Urban schools generate large quantities of organic waste that typically go to landfills. This initiative sets up compost pits and educates students about sustainable waste management. Children actively participate in maintaining the system, building long-term eco-conscious habits from a young age.',
  (SELECT id FROM cause_domains WHERE name = 'Education'),
  (SELECT id FROM cause_aid_types WHERE name = 'Environmental Support'),
  60000.00, 8000.00,
  '2027-02-28 23:59:59+05:30',
  'https://example.com/covers/compost.jpg'
),
-- 20 Disaster Relief Assistance
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Drought Relief: Water Tanker & Bore Repairs',
  'Several villages are experiencing severe drought, resulting in acute water scarcity for drinking and agriculture. This project dispatches water tankers and repairs community borewells to ensure consistent supply. It aims to prevent migration and support vulnerable households during climate emergencies.',
  (SELECT id FROM cause_domains WHERE name = 'Urgent'),
  (SELECT id FROM cause_aid_types WHERE name = 'Disaster Relief Assistance'),
  300000.00, 60000.00,
  '2026-10-31 23:59:59+05:30',
  'https://example.com/covers/drought_relief.jpg'
),
-- 21 Monetary Donations
(
  (SELECT id FROM organizations WHERE organization_name = 'Nari Utthan Foundation'),
  'Scholarships for Women in Tech',
  'Technically talented girls from rural and tier-2 cities often discontinue higher education due to lack of financial support. This initiative offers merit-based scholarships to women pursuing technology degrees, helping them enter high-growth career fields. It also provides mentorship from industry professionals.',
  (SELECT id FROM cause_domains WHERE name = 'Women'),
  (SELECT id FROM cause_aid_types WHERE name = 'Monetary Donations'),
  350000.00, 90000.00,
  '2027-07-31 23:59:59+05:30',
  'https://example.com/covers/scholarships.jpg'
),
-- 22 Medical Assistance
(
  (SELECT id FROM organizations WHERE organization_name = 'Jal Raksha Initiative'),
  'Clean Water & Sanitation Health Drive',
  'Unsafe water sources and poor sanitation contribute to deadly diseases in rural areas. This program combines water purification device distribution, hygiene workshops, and primary health screenings. The goal is to improve long-term health outcomes in water-stressed communities.',
  (SELECT id FROM cause_domains WHERE name = 'Health'),
  (SELECT id FROM cause_aid_types WHERE name = 'Medical Assistance'),
  260000.00, 32000.00,
  '2027-01-31 23:59:59+05:30',
  'https://example.com/covers/wash_health.jpg'
),
-- 23 Goods & Resources for Children
(
  (SELECT id FROM organizations WHERE organization_name = 'Care for Strays Collective'),
  'School Stationery Packs for Underprivileged Kids',
  'Children from low-income municipal school families often attend classes without basic learning tools. This initiative distributes essential stationery kits, including books, bags, pencils, and uniforms, to enhance classroom readiness. It ensures access to dignity in education from an early age.',
  (SELECT id FROM cause_domains WHERE name = 'Children'),
  (SELECT id FROM cause_aid_types WHERE name = 'Goods & Resources'),
  50000.00, 7000.00,
  '2026-10-15 23:59:59+05:30',
  'https://example.com/covers/stationery.jpg'
),
-- 24 Arts & Culture volunteering
(
  (SELECT id FROM organizations WHERE organization_name = 'Green Earth India'),
  'Heritage Wall Mural Project',
  'Local heritage sites are losing cultural identity due to neglect and urbanisation. This project brings together volunteer artists and youth to paint significant public murals that revive cultural awareness. It also supports local tourism and beautification of shared community spaces.',
  (SELECT id FROM cause_domains WHERE name = 'Arts & Culture'),
  (SELECT id FROM cause_aid_types WHERE name = 'Volunteering'),
  45000.00, 6000.00,
  '2027-04-30 23:59:59+05:30',
  'https://example.com/covers/mural.jpg'
);
