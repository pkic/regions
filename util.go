package regions

import (
	"regexp"
	"strings"
)

var (
	// Some values contain meta data that we don't want
	metaDataFormat   = regexp.MustCompile(`([\'|\,|\s]$|\*|\([^\)]*\)|\[[^\]]*\]|[-|,|\s]+([A-Z]\.)+)+`)
	doubleWhitespace = regexp.MustCompile(`\s{2,}`)

	// Initially generated with Google Translate with English as the source language
	// - Replaced ) with )
	// - French (fr) [Added 'en']
	// - Polish (pl) [Added 'Prowincja']
	// - Portuguese (pt) [Added 'do']
	// - Romanian (ro) [Added 'a']
	// - Slovenian (sl) [Changed Province to Provinca]
	// - China (zh) [Added 'Sheng', 'Shi', and 'SAR']
	// - Corsican (co) [Removed space in 'l ']
	// - Russian (ru) [Added 'oblast']
	localPreOrSuffix = []*regexp.Regexp{
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Region|Capital city|Capital region|Community|Autonomous|Republic|City|Territory|of|the)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // English (en)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincie|Regio|Hoofdstad|Hoofdstedelijk gewest|Gemeenschap|autonoom|Republiek|stad|Gebied|van|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Dutch (nl)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((provinsie|streek|Hoofstad|streek kapitaal|Gemeenskap|outonome|Republiek|Stad|gebied|van|die)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                // Afrikaans (af)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((krahinë|Rajon|Kryeqytet|rajonin e kryeqytetit|Community|autonom|republikë|qytet|territor|të)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                // Albanian (sq)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ክፍለ ሀገር|ክልል|ዋና ከተማ|ካፒታል ክልል|ኅብረተሰብ|ነጻ|ሬፑብሊክ|ከተማ|ግዛት|የ|የ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                    // Amharic (am)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((المحافظة|منطقة|العاصمة|منطقة العاصمة|تواصل اجتماعي|واثق من نفسه|جمهورية|مدينة|منطقة|من|ال)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Arabic (ar)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((նահանգ|մարզ|Մայրաքաղաք|մայրաքաղաքը մարզ|Համայնք|ինքնավար|հանրապետություն|քաղաք|տարածքը|Հյուրատետր|որ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                       // Armenian (hy)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((vilayət|rayon|Paytaxt|Capital region|icma|muxtar|respublika|şəhər|ərazi|of|the)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                             // Azerbaijani (az)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Probintzia|Region|Hiriburua|Capital eskualdean|Komunitatea|Autonomia|Errepublika|hiria|Lurraldea|of|du)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                     // Basque (eu)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((правінцыя|вобласць|Сталіца|сталічны рэгіён|супольнасць|аўтаномны|рэспубліка|горад|тэрыторыя|з)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Belarusian (be)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((প্রদেশ|অঞ্চল|রাজধানী শহর|ক্যাপিটাল অঞ্চল|সম্প্রদায়|স্বায়ত্তশাসিত|প্রজাতন্ত্র|শহর|এলাকা|এর|দ্য)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Bengali (Bangla) (bn)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincija|regija|Glavni grad|Capital regija|Zajednica|autonoman|republika|grad|teritorija|od|u)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Bosnian (bs)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провинция|област|Столица|област Капитал|общност|автономен|република|град|Територия|на|на)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                   // Bulgarian (bg)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((နယ်ပယ်|ဒေသ|မြို့တော်|မြို့တော်ဒေသ|အဝန်း|ကိုယ်ပိုင်အုပ်ချုပ်ရေးရသော|သမတနိုင်ငံ|မြို့|နယ်မြေတွေကို|၏|အဆိုပါ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                  // Burmese (my)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((província|regió|Ciutat capital|regió de la capital|comunitat|autònom|República|ciutat|territori|de|la)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                      // Catalan (ca)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Region|likulu|Capital dera|Community|Autonomous|Republic|City|gawo|wa|ndi)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                         // Chichewa, Chewa, Nyanja (ny)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((省|地区|首都|首都地区|社区|自主性|共和国|城市|领土|的|这|Sheng|Shi|SAR)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                            // Chinese (zh)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((省|地区|首都|首都地区|社区|自主性|共和国|城市|领土|的|这)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                                          // Chinese (Simplified) (zh-Hans)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((省|地區|首都|首都地區|社區|自主性|共和國|城市|領土|的|這)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                                          // Chinese (Traditional) (zh-Hant)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Pruvincia|Region|a cità capitale|righjoni Capital|cumunità|autònuma|Republic|cità|Territory|di|l')(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Corsican (co)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Pokrajina|Regija|Glavni grad|Capital regija|Zajednica|autonoman|Republika|Grad|Teritorija|od)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                               // Croatian (hr)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincie|Kraj|Hlavní město|Capital region|Společenství|Autonomní|Republika|Město|Území|z)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Czech (cs)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provins|Område|Hovedstad|region hovedstaden|Fællesskab|Autonom|Republik|by|Territorium|af|det)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Danish (da)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provinco|Regiono|Ĉefurbo|ĉefurbo regiono|komunumo|aŭtonoma|Respubliko|Urbo|teritorio|el|la)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                 // Esperanto (eo)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((provints|piirkond|Pealinn|Capital piirkonnas|kogukond|autonoomne|Vabariik|linn|territoorium|kohta)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Estonian (et)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((maakunta|alue|Pääkaupunki|pääkaupunkiseutu|Yhteisö|Autonominen|Tasavalta|Kaupunki|Alue|of)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Finnish (fi)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Région|Capitale|région de la capitale|Communauté|Autonome|République|Ville|Territoire|de|les|en)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                   // French (fr)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((provincia|rexión|capital|rexión da capital|comunidade|autónomo|República|cidade|territorio|de|o)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Galician (gl)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Mòr-roinn|Region|prìomh-bhaile|calpa roinn|Coimhearsnachd|Neo-eisimeileach Thìr|Poblachd|City|Territory|de|a ')(?:\A|\z|\s|\.|\,|\;)+)+)+`),             // Gaelic (Scottish) (gd)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((პროვინციაში|რეგიონი|Დედაქალაქი|დედაქალაქის რეგიონი|Community|Ავტონომიური|რესპუბლიკა|ქალაქი|ტერიტორია|საქართველოს)(?:\A|\z|\s|\.|\,|\;)+)+)+`),           // Georgian (ka)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provinz|Region|Hauptstadt|Hauptstadtregion|Gemeinschaft|autonom|Republik|Stadt|Gebiet|von|das)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // German (de)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Επαρχία|Περιοχή|Πρωτεύουσα|περιοχή της πρωτεύουσας|Κοινότητα|Αυτονόμος|Δημοκρατία|Πόλη|Εδαφος|του|ο)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                        // Greek (el)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((પ્રાંત|પ્રદેશ|રાજધાની શહેર|કેપિટલ પ્રદેશ|સમુદાય|સ્વાયત્ત|રિપબ્લિક|શહેરનું|ટેરિટરી|ના|આ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                     // Gujarati (gu)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Pwovens|rejyon|Kapital vil|rejyon Kapital|Kominote|otonòm|Repiblik|City|Teritwa|nan|nan)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                    // Haitian Creole (ht)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((lardin|yankin|babban birnin|babban birnin yankin|Community|mai cin gashin kanta|Jamhuriyar|City|Territory|na|da)(?:\A|\z|\s|\.|\,|\;)+)+)+`),            // Hausa (ha)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((מָחוֹז|אזור|עיר בירה|אזור קפיטל|הקהילה|אוטונומי|רפובליקה|עִיר|שֶׁטַח|שֶׁל|ה)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                // Hebrew (he)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((प्रांत|क्षेत्र|राजधानी|राजधानी क्षेत्र|समुदाय|स्वायत्तशासी|गणतंत्र|Faridabad|क्षेत्र|का)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                    // Hindi (hi)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Tartomány|Vidék|Főváros|Főváros régióban|Közösség|Autonóm|Köztársaság|Város|Terület|nak,-nek|a)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                             // Hungarian (hu)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Region|Höfuðborg|Capital Region|Community|Autonomous|Republic|Borg|Territory|af|sem)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                               // Icelandic (is)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Region|Capital obodo|Capital region|Community|Autonomous|Republic|City|Territory|nke|na)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                           // Igbo (ig)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Cúige|Réigiún|Príomhchathair|réigiún caipitil|Pobail|Autonomous|Poblacht|Cathair|Críoch|de|an)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Irish (ga)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincia|Regione|Capitale|regione della capitale|Comunità|Autonomo|Repubblica|Città|Territorio|di|il)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                      // Italian (it)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((州|領域|首都|首都圏|コミュニティ|自主的な|共和国|市|地域|の|インクルード)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                                  // Japanese (ja)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Langkawi|Region|kutha Capital|wilayah Capital|komunitas|otonomi|Republik|City|Territory|saka|ing)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                           // Javanese (jv)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ಪ್ರಾಂತ್ಯ|ಪ್ರದೇಶ|ರಾಜಧಾನಿ|ರಾಜಧಾನಿ ಪ್ರದೇಶ|ಸಮುದಾಯ|ಸ್ವಾಯತ್ತ|ರಿಪಬ್ಲಿಕ್|ಸಿಟಿ|ಟೆರಿಟರಿ|ಆಫ್|ದಿ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                     // Kannada (kn)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провинция|аймақ|Астана|капитал облысы|қауым|автономиялық|республика|қала|аумақ|туралы|The)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Kazakh (kk)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ខេត្ត|តំបន់|រាជធានី|តំបន់រដ្ឋធានី|សហគមន៍|ស្វយ័ត|សាធារណរដ្ឋ|ទីក្រុង|ដែនដី|នៃ|នេះ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                            // Khmer (km)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Intara|Region|umurwa mukuru|Capital karere|Community|Ubwiyobore|Rusaye|Umugi|Territory|ya|mu)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                               // Kinyarwanda (Rwanda) (rw)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Облус|регион|борбор шаар|Капитал аймак|Community|автономдуу|республика|Сити|территория|боюнча|жана)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                         // Kyrgyz (ky)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((지방|부위|자본 도시|수도권|커뮤니티|자발적인|공화국|시티|영토|의|그만큼)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                                                  // Korean (ko)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Herêm|Herêm|Paytext|herêmê Capital|Civatî|Serbixwe|Cumhurîyet|Bajar|Herêm|ji|ew)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                            // Kurdish (ku)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ແຂວງ|ພາກພື້ນ|ນະຄອນຫຼວງ|ເຂດນະຄອນຫຼວງ|ຊຸມຊົນ|ເອກະລາດ|ສາທາລະນະ|ເມືອງ|ອານາເຂດຂອງ|ຂອງ|ໄດ້)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                       // Lao (lo)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|regio|Oppida prima|regionem capitis|Community|sui iuris|publica|Urbs|Territorium|autem|quod)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                       // Latin (la)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((province|Novads|Galvaspilsēta|Capital reģions|kopiena|autonoms|republika|pilsēta|teritorija|no)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                             // Latvian (Lettish) (lv)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((provincija|regionas|Sostinė|Capital Region|bendruomenė|savarankiškas|respublika|miestas|teritorija|apie)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                    // Lithuanian (lt)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provënz|Regioun|Haaptstad|Capital Regioun|Communautéit|autonom|Republik|City|Territoire|vun|der)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Luxembourgish (lb)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провинција|регионот|Главен град|капитал регион|заедница|автономна|Република|Сити|територија|на|на)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Macedonian (mk)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((-tokony eran'ny fanjakana|Region|Renivohitra|Capital faritra|fiaraha-monina|Mizaka tena|Republic|City|FARITANY|ny|ny)(?:\A|\z|\s|\.|\,|\;)+)+)+`),       // Malagasy (mg)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((wilayah|rantau|Ibu negeri|Capital Region|komuniti|autonomi|Republic|City|wilayah|daripada|yang)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                             // Malay (ms)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((പ്രവിശ്യ|പ്രദേശം|തലസ്ഥാന നഗരം|തലസ്ഥാന|സമൂഹം|സയംശാസിതമായ|ജനാധിപത്യഭരണം|നഗരം|ടെറിട്ടറി|എന്ന|The)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Malayalam (ml)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((provinċja|reġjun|Belt kapitali|reġjun kapitali|Komunità|awtonoma|Repubblika|belt|territorju|ta|il)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Maltese (mt)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Porowini|Region|Capital pa|Capital rohe|hapori|motuhake|Republic|City|Territory|o|te)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                       // Maori (mi)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((प्रांत|प्रदेश|राजधानी|कॅपिटल प्रदेश|समुदाय|स्वायत्त|प्रजासत्ताक|सिटी|प्रदेश|च्या|अगोदर निर्देश केलेल्या बाबीसंबंधी बोलताना)(?:\A|\z|\s|\.|\,|\;)+)+)+`), // Marathi (mr)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincie|Regiune|Capitala|regiunea de capital|Comunitate|Autonom|Republică|Oraș|Teritoriu|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Moldavian (mo)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((аймгийн|бүс нутаг|Нийслэл хот|Нийслэлийн бүс|Олон нийтийн|Өөртөө Засах|Бүгд Найрамдах Улс|хот|газар нутаг|нь|The)(?:\A|\z|\s|\.|\,|\;)+)+)+`),           // Mongolian (mn)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((प्रान्त|क्षेत्र|राजधानी|राजधानी क्षेत्र|समुदाय|स्वायत्त|गणतन्त्र|शहर|इलाका|को|को)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                           // Nepali (ne)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provins|Region|Hovedstad|Capital region|Samfunnet|Autonomous|Republikk|By|Territory|av|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Norwegian (no)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provins|Region|Hovedstad|Capital region|Samfunnet|Autonomous|Republikk|By|Territory|av|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Norwegian bokmål (nb)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ପ୍ରେଦଶ|ଅଞ୍ଚଳ|ରାଜଧାନୀ ସହର|Capital ଅଞ୍ଚଳ|ସମ୍ପ୍ରଦାୟ|ସ୍ବୟଂ ଶାସିତ|ଗଣତନ୍ତ୍ର|ସହର|ଭୂଭାଗ|ର|େଯମାେନ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                   // Oriya (or)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ولایت|سیمه|پلازمېنه ښار|پلازمیینه سیمه|د ټولنې|خپلواکه|جمهوریت|ښار|خاوره|د|د)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                               // Pashto, Pushto (ps)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((استان|منطقه|پایتخت|منطقه پایتخت|انجمن|خود مختار|جمهوری|شهرستان|قلمرو|از)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                    // Persian (Farsi) (fa)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Województwo|Region|Stolica|Region stołeczny|Społeczność|Autonomiczny|Republika|Miasto|Terytorium|z|Prowincja)(?:\A|\z|\s|\.|\,|\;)+)+)+`),               // Polish (pl)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Província|Região|Capital|região da capital|Comunidade|Autônomo|República|Cidade|Território|de|a|do)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                         // Portuguese (pt)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ਸੂਬੇ|ਖੇਤਰ|ਰਾਜਧਾਨੀ|ਰਾਜਧਾਨੀ ਖੇਤਰ|ਕਮਿਊਨਿਟੀ|ਆਟੋਨੋਮਸ|ਗਣਤੰਤਰ|ਸਿਟੀ|ਟੈਰੀਟਰੀ|ਦੇ|ਇਹ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                  // Punjabi (Eastern) (pa)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincie|Regiune|Capitala|regiunea de capital|Comunitate|Autonom|Republică|Oraș|Teritoriu|de|a)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Romanian (ro)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Провинция|Область, край|Столица|Столичный регион|Сообщество|автономный|республика|Город|территория|из|oblast)(?:\A|\z|\s|\.|\,|\;)+)+)+`),               // Russian (ru)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((setate|vāega|laumua|Tupe Faavae itulagi|Community|tū toʻatasi|mālō faiperesitene|aʻai|atunuʻu|a|le)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                         // Samoan (sm)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провинција|регија|Главни град|Цапитал region|заједница|аутономан|република|град|територија|од)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                              // Serbian (sr)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|sebakeng|motse-moholo|motse-moholo oa sebaka|Community|ikemetseng|Republic|City|Territory|ya|ho)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                   // Sesotho (st)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((dunhu|nharaunda|Guta guru|Capital nharaunda|Community|zvionera|utongi hwegutsaruzhinji|guta|nzvimbo|pamusoro|ari)(?:\A|\z|\s|\.|\,|\;)+)+)+`),           // Shona (sn)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((صوبي|علائقي|گادي جو شھر|گاديء واري علائقي|ڪميونٽي|خودمختيار|جمهوريه|شهر|سڏجي ٿو|جي|جي)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                      // Sindhi (sd)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((පළාත|කලාපයේ|අග නගරය|අගනුවර කලාපයේ|ප්රජා|ස්වාධීන|ජනරජය|නගරය|භූමිය|වල|එම)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                     // Sinhalese (si)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincie|kraj|Hlavné mesto|capital región|spoločenstvo|autonómne|republika|veľkomesto|územie|z)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Slovak (sk)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provinca|regija|Glavno mesto|Capital regija|Skupnosti|avtonomna|republika|Kraj|ozemlje|za)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Slovenian (sl)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Gobolka|gobolka|magaalada Capital|gobolka Capital|bulshada|Banaan|Republic|City|Territory|of|ah)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Somali (so)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provincia|Región|Ciudad capital|región de la capital|Comunidad|Autónomo|República|Ciudad|Territorio|de|la)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                  // Spanish (es)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Propinsi|daerah|kota Capital|wewengkon ibukota|masarakat|otonomi|republik|kota|wewengkon|ti|éta)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Sundanese (su)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Mkoa|Mkoa|Mji mkuu|Capital kanda|Community|Autonomous|Jamhuri|City|Wilaya|ya|the)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                           // Swahili (Kiswahili) (sw)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provins|Område|Huvudstad|kapital region|gemenskap|Autonom|republik|Stad|Territorium|av|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Swedish (sv)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((lalawigan|Rehiyon|capital city|Capital region|komunidad|nagsasarili|republika|lungsod|lupain|ng|ang)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                        // Tagalog (tl)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((вилоят|вилоят|Пойтахт|минтақаи Пойтахт|Community|худмухторӣ|Ҷумҳурии|шаҳр|қаламрави|аз|ба)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Tajik (tg)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((மாகாணம்|பகுதி|தலை நாகரம்|தலைநகர பிராந்தியம்|சமூக|தன்னாட்சிப்|குடியரசு|நகரம்|பிரதேசம்|இன்|தி)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                // Tamil (ta)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провинция|өлкә|Башкала|капитал районы|Җәмгыятьтәге|үзидарәле|республика|шәһәр|территория|һәм|бу)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Tatar (tt)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((ప్రావిన్స్|ప్రాంతం|రాజధాని నగరం|రాజధాని ప్రాంతం|సంఘం|అటానమస్|రిపబ్లిక్|నగరం|భూభాగం|ఆఫ్|ది)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Telugu (te)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((จังหวัด|ภูมิภาค|เมืองหลวง|ภูมิภาคทุน|ชุมชน|อิสระ|สาธารณรัฐ|เมือง|อาณาเขต|ของ)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                               // Thai (th)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Bölge|bölge|Başkent|Sermaye bölge|Topluluk|özerk|cumhuriyet|Kent|bölge|nın-nin)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                             // Turkish (tr)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((welaýat|sebit|Paýtagt|Capital sebit|jemgyýetçilik|özbaşdak|respublika|şäher|meýdany|we|mälimlik görkeziji artikl)(?:\A|\z|\s|\.|\,|\;)+)+)+`),           // Turkmen (tk)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((өлкә|رايون|پايتەخت شەھەر|پايتەخت رايون|مەھەللە|ئاپتونوم|жумһүрийәт|شەھەر|территория|نىڭ|ишлитилмәйду)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                       // Uyghur (ug)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((провінція|область|Столиця|столичний регіон|співтовариство|автономний|республіка|місто|територія|з)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                          // Ukrainian (uk)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((صوبہ|علاقہ|دارالحکومت|کیپٹل خطے|برادری|خود مختار|جمہوریہ|شہر|علاقہ|کے)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                                      // Urdu (ur)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((viloyat|mintaqa|Poytaxt shahar|Capital viloyati|jamoa|Avtonom|respublika|shahar|hududi|ning|The)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                            // Uzbek (uz)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Tỉnh|Khu vực|Thủ đô|khu vực vốn|cộng đồng|tự trị|nước cộng hòa|thành phố|lãnh thổ|của|các)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                  // Vietnamese (vi)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Talaith|rhanbarth|Prifddinas|rhanbarth y Brifddinas|cymunedol|ymreolaethol|Republic|City|Tiriogaeth|o|y)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                    // Welsh (cy)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Provinsje|Regio|Haadstêd|Capital regio|Mienskip|Autonoom|Republyk|Stêd|Gebiet|fan|de)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                       // Western Frisian (fy)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((kwiPhondo|Region|Isixeko esikhulu|indawo Capital|Community|lizimele|Republic|City|Territory|of|i)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                           // Xhosa (xh)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Province|Region|Olú ìlú|Capital ekun|Community|adase|Republic|City|agbegbe|ti|awọn)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                                         // Yoruba (yo)
		regexp.MustCompile(`(?i)((?:\A|\z|\s|\.|\,|\;)+((Isifundazwe|Isifunda|Inhloko-dolobha|Capital esifundeni|Umphakathi|Autonomous|Republic|City|Territory|ka|le)(?:\A|\z|\s|\.|\,|\;)+)+)+`),                // Zulu (zu)
	}
)

func removeMetaData(name string) string {
	name = metaDataFormat.ReplaceAllLiteralString(name, "")
	for _, l := range localPreOrSuffix {
		name = l.ReplaceAllLiteralString(name, " ")
	}
	name = doubleWhitespace.ReplaceAllLiteralString(name, " ")
	return strings.TrimSpace(name)
}

func stringInSlice(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, str) {
			return true
		}
	}
	return false
}
