package ppos_test

const SystemContractCode = "0x6080604052600436106101e25760003560e01c80638a11d7c911610102578063b6c8851911610095578063d6c0edad11610064578063d6c0edad14610e8d578063db78dd2814610e95578063efd8d8e214610e95578063f2888dbb14610ec7576101e2565b8063b6c8851914610b59578063be64569214610e1e578063c253c38414610e33578063c967f90f14610e61576101e2565b8063a224cee7116100d1578063a224cee7146106bb578063a406fcb714610736578063a43569b314610901578063afeea11514610b44576101e2565b80638a11d7c9146105495780638b0e9f3f1461060e57806398e3b626146106235780639de7025814610656576101e2565b80634b3d500b1161017a5780636846992a116101495780636846992a146103e35780636969a25c146104935780637f4f95fa146104bd57806382bd3d9214610516576101e2565b80634b3d500b1461034a578063588e9178146103745780635dd095901461039b5780636233be5d146103ce576101e2565b806326476204116101b657806326476204146102a75780633a061bd3146102cd57806340550a1c146102e257806340a141ff14610315576101e2565b8062362a77146101e7578063158ef93e1461022e5780631b5e358c14610243578063222d3b0514610274575b600080fd5b3480156101f357600080fd5b5061021a6004803603602081101561020a57600080fd5b50356001600160a01b0316610efa565b604080519115158252519081900360200190f35b34801561023a57600080fd5b5061021a611161565b34801561024f57600080fd5b5061025861116a565b604080516001600160a01b039092168252519081900360200190f35b34801561028057600080fd5b5061021a6004803603602081101561029757600080fd5b50356001600160a01b0316611170565b61021a600480360360208110156102bd57600080fd5b50356001600160a01b0316611393565b3480156102d957600080fd5b506102586116f2565b3480156102ee57600080fd5b5061021a6004803603602081101561030557600080fd5b50356001600160a01b03166116f8565b34801561032157600080fd5b506103486004803603602081101561033857600080fd5b50356001600160a01b0316611753565b005b34801561035657600080fd5b506102586004803603602081101561036d57600080fd5b503561189b565b34801561038057600080fd5b506103896118c2565b60408051918252519081900360200190f35b3480156103a757600080fd5b50610348600480360360208110156103be57600080fd5b50356001600160a01b03166118c8565b3480156103da57600080fd5b50610258611921565b3480156103ef57600080fd5b506103486004803603604081101561040657600080fd5b810190602081018135600160201b81111561042057600080fd5b82018360208201111561043257600080fd5b803590602001918460208302840111600160201b8311171561045357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250611927915050565b34801561049f57600080fd5b50610258600480360360208110156104b657600080fd5b5035611b6e565b3480156104c957600080fd5b506104f8600480360360408110156104e057600080fd5b506001600160a01b0381358116916020013516611b7b565b60408051938452602084019290925282820152519081900360600190f35b34801561052257600080fd5b5061021a6004803603602081101561053957600080fd5b50356001600160a01b0316611bb8565b34801561055557600080fd5b5061057c6004803603602081101561056c57600080fd5b50356001600160a01b0316611e1c565b60405180886001600160a01b0316815260200187600481111561059b57fe5b815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156105f45781810151838201526020016105dc565b505050509050019850505050505050505060405180910390f35b34801561061a57600080fd5b50610389612262565b34801561062f57600080fd5b5061021a6004803603602081101561064657600080fd5b50356001600160a01b0316612268565b34801561066257600080fd5b5061066b6122ba565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156106a757818101518382015260200161068f565b505050509050019250505060405180910390f35b3480156106c757600080fd5b50610348600480360360208110156106de57600080fd5b810190602081018135600160201b8111156106f857600080fd5b82018360208201111561070a57600080fd5b803590602001918460208302840111600160201b8311171561072b57600080fd5b50909250905061231c565b34801561074257600080fd5b5061021a600480360360c081101561075957600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561078357600080fd5b82018360208201111561079557600080fd5b803590602001918460018302840111600160201b831117156107b657600080fd5b919390929091602081019035600160201b8111156107d357600080fd5b8201836020820111156107e557600080fd5b803590602001918460018302840111600160201b8311171561080657600080fd5b919390929091602081019035600160201b81111561082357600080fd5b82018360208201111561083557600080fd5b803590602001918460018302840111600160201b8311171561085657600080fd5b919390929091602081019035600160201b81111561087357600080fd5b82018360208201111561088557600080fd5b803590602001918460018302840111600160201b831117156108a657600080fd5b919390929091602081019035600160201b8111156108c357600080fd5b8201836020820111156108d557600080fd5b803590602001918460018302840111600160201b831117156108f657600080fd5b50909250905061267d565b34801561090d57600080fd5b506109346004803603602081101561092457600080fd5b50356001600160a01b0316612c89565b60405180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019080838360005b83811015610981578181015183820152602001610969565b50505050905090810190601f1680156109ae5780820380516001836020036101000a031916815260200191505b5086810385528a5181528a516020918201918c019080838360005b838110156109e15781810151838201526020016109c9565b50505050905090810190601f168015610a0e5780820380516001836020036101000a031916815260200191505b5086810384528951815289516020918201918b019080838360005b83811015610a41578181015183820152602001610a29565b50505050905090810190601f168015610a6e5780820380516001836020036101000a031916815260200191505b5086810383528851815288516020918201918a019080838360005b83811015610aa1578181015183820152602001610a89565b50505050905090810190601f168015610ace5780820380516001836020036101000a031916815260200191505b50868103825287518152875160209182019189019080838360005b83811015610b01578181015183820152602001610ae9565b50505050905090810190601f168015610b2e5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b348015610b5057600080fd5b5061066b6130bd565b348015610b6557600080fd5b5061021a600480360360a0811015610b7c57600080fd5b810190602081018135600160201b811115610b9657600080fd5b820183602082011115610ba857600080fd5b803590602001918460018302840111600160201b83111715610bc957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610c1b57600080fd5b820183602082011115610c2d57600080fd5b803590602001918460018302840111600160201b83111715610c4e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610ca057600080fd5b820183602082011115610cb257600080fd5b803590602001918460018302840111600160201b83111715610cd357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610d2557600080fd5b820183602082011115610d3757600080fd5b803590602001918460018302840111600160201b83111715610d5857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610daa57600080fd5b820183602082011115610dbc57600080fd5b803590602001918460018302840111600160201b83111715610ddd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061311d945050505050565b348015610e2a57600080fd5b506103896132c2565b348015610e3f57600080fd5b50610e486132ca565b6040805192835260208301919091528051918290030190f35b348015610e6d57600080fd5b50610e766132df565b6040805161ffff9092168252519081900360200190f35b6103486132e4565b348015610ea157600080fd5b50610eaa613484565b6040805167ffffffffffffffff9092168252519081900360200190f35b348015610ed357600080fd5b5061021a60048036036020811015610eea57600080fd5b50356001600160a01b0316613489565b600033816001600160a01b038416600090815260016020526040902054600160a01b900460ff166004811115610f2c57fe5b1415610f75576040805162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b6001600160a01b03838116600090815260016020526040902054811690821614610fd05760405162461bcd60e51b815260040180806020018281038252602e815260200180614855602e913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902060090154436064909101111561102f5760405162461bcd60e51b815260040180806020018281038252605c815260200180614700605c913960600191505060405180910390fd5b6001600160a01b0383166000908152600160205260409020600701548061109d576040805162461bcd60e51b815260206004820152601a60248201527f596f7520646f6e2774206861766520616e792070726f66697473000000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526001602052604081206007810191909155436009909101558015611102576040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015611100573d6000803e3d6000fd5b505b816001600160a01b0316846001600160a01b03167f51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb98342604051808381526020018281526020019250505060405180910390a36001925050505b919050565b60005460ff1681565b61f00181565b3360008181526002602090815260408083206001600160a01b0386168452825280832060019092528220549192918390600160a01b900460ff1660048111156111b557fe5b14156111fe576040805162461bcd60e51b81526020600482015260136024820152721d985b1a59185d1bdc881b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b6001810154611254576040805162461bcd60e51b815260206004820152601960248201527f596f75206861766520746f20756e7374616b6520666972737400000000000000604482015290519081900360640190fd5b43606467ffffffffffffffff1682600101540111156112a45760405162461bcd60e51b815260040180806020018281038252602181526020018061475c6021913960400191505060405180910390fd5b80546112f2576040805162461bcd60e51b8152602060048201526018602482015277596f7520646f6e2774206861766520616e79207374616b6560401b604482015290519081900360640190fd5b80546000808355600183018190556040516001600160a01b0385169183156108fc02918491818181858888f19350505050158015611334573d6000803e3d6000fd5b50846001600160a01b0316836001600160a01b03167fa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c149556198342604051808381526020018281526020019250505060405180910390a3506001949350505050565b6000805460ff166113da576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b333460016001600160a01b038516600090815260016020526040902054600160a01b900460ff16600481111561140c57fe5b1480611445575060026001600160a01b038516600090815260016020526040902054600160a01b900460ff16600481111561144357fe5b145b6114805760405162461bcd60e51b815260040180806020018281038252602d8152602001806147e4602d913960400191505060405180910390fd5b6001600160a01b03808316600090815260026020908152604080832093881683529290522060010154156114e55760405162461bcd60e51b81526004018080602001828103825260228152602001806147c26022913960400191505060405180910390fd5b6001600160a01b0384166000908152600160208190526040909120908101546305f5e1009061151490846138c5565b1015611567576040805162461bcd60e51b815260206004820152601860248201527f5374616b696e6720636f696e73206e6f7420656e6f7567680000000000000000604482015290519081900360640190fd5b6001600160a01b038084166000908152600260209081526040808320938916835292905220546115e057600a810180546001600160a01b038086166000818152600260208181526040808420958d1684529481529382200184905560018401855593845292200180546001600160a01b03191690911790555b60018101546115ef90836138c5565b600182015560028154600160a01b900460ff16600481111561160d57fe5b1461162457805460ff60a01b1916600160a11b1781555b6116378582600101548360070154613928565b6001600160a01b0380841660009081526002602090815260408083209389168352929052205461166790836138c5565b6001600160a01b038085166000908152600260209081526040808320938a168352929052205560055461169a90836138c5565b6005556040805183815242602082015281516001600160a01b0380891693908716927fb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075929081900390910190a3506001949350505050565b61f00081565b6000805b60035481101561174a57826001600160a01b03166003828154811061171d57fe5b6000918252602090912001546001600160a01b0316141561174257600191505061115c565b6001016116fc565b50600092915050565b3361f001146117a0576040805162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015290519081900360640190fd5b6001600160a01b0381166000908152600160205260409020600701546117c582613cc7565b60045460011015611897576117d982613ddf565b600754604080516315ea278160e01b81526001600160a01b038581166004830152915191909216916315ea27819160248083019260209291908290030181600087803b15801561182857600080fd5b505af115801561183c573d6000803e3d6000fd5b505050506040513d602081101561185257600080fd5b50506040805182815242602082015281516001600160a01b038516927fa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe928290030190a25b5050565b600481815481106118a857fe5b6000918252602090912001546001600160a01b0316905081565b60065481565b3361f00114611915576040805162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015290519081900360640190fd5b61191e81613cc7565b50565b61f00281565b334114611968576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b4360009081526009602090815260408083206001845290915290205460ff16156119d9576040805162461bcd60e51b815260206004820152601a60248201527f56616c696461746f727320616c72656164792075706461746564000000000000604482015290519081900360640190fd5b60005460ff16611a1f576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b80804381611a2957fe5b0615611a6f576040805162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015290519081900360640190fd5b43600090815260096020908152604080832060018085529252909120805460ff191690911790558251611ae0576040805162461bcd60e51b815260206004820152601460248201527356616c696461746f722073657420656d7074792160601b604482015290519081900360640190fd5b8251611af3906003906020860190614571565b507feacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5836040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611b56578181015183820152602001611b3e565b505050509050019250505060405180910390a1505050565b600381815481106118a857fe5b6001600160a01b03918216600090815260026020818152604080842094909516835292909252919091208054600182015491909201549192909190565b60003361f00214611c09576040805162461bcd60e51b815260206004820152601660248201527550726f706f73616c20636f6e7472616374206f6e6c7960501b604482015290519081900360640190fd5b60005460ff16611c4f576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b60036001600160a01b038316600090815260016020526040902054600160a01b900460ff166004811115611c7f57fe5b14158015611cbb575060046001600160a01b038316600090815260016020526040902054600160a01b900460ff166004811115611cb857fe5b14155b15611cc85750600161115c565b60046001600160a01b038316600090815260016020526040902054600160a01b900460ff166004811115611cf857fe5b1415611db957600854604080516363e1d45160e01b81526001600160a01b038581166004830152915191909216916363e1d4519160248083019260209291908290030181600087803b158015611d4d57600080fd5b505af1158015611d61573d6000803e3d6000fd5b505050506040513d6020811015611d7757600080fd5b5051611db9576040805162461bcd60e51b815260206004820152600c60248201526b18db19585b8819985a5b195960a21b604482015290519081900360640190fd5b6001600160a01b038216600081815260016020908152604091829020805460ff60a01b1916600160a01b179055815142815291517fd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb4009281900390910190a2919050565b6000806000806000806060611e2f6145d6565b6001600160a01b038981166000908152600160209081526040918290208251610100810190935280549384168352919290830190600160a01b900460ff166004811115611e7857fe5b6004811115611e8357fe5b815260200160018201548152602001600282016040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f3a5780601f10611f0f57610100808354040283529160200191611f3a565b820191906000526020600020905b815481529060010190602001808311611f1d57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611fdc5780601f10611fb157610100808354040283529160200191611fdc565b820191906000526020600020905b815481529060010190602001808311611fbf57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561206e5780601f106120435761010080835404028352916020019161206e565b820191906000526020600020905b81548152906001019060200180831161205157829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156121025780601f106120d757610100808354040283529160200191612102565b820191906000526020600020905b8154815290600101906020018083116120e557829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156121965780601f1061216b57610100808354040283529160200191612196565b820191906000526020600020905b81548152906001019060200180831161217957829003601f168201915b5050505050815250508152602001600782015481526020016008820154815260200160098201548152602001600a820180548060200260200160405190810160405280929190818152602001828054801561221a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116121fc575b505050505081525050905080600001518160200151826040015183608001518460a001518560c001518660e00151975097509750975097509750975050919395979092949650565b60055481565b6000805b60045481101561174a57826001600160a01b03166004828154811061228d57fe5b6000918252602090912001546001600160a01b031614156122b257600191505061115c565b60010161226c565b6060600380548060200260200160405190810160405280929190818152602001828054801561231257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122f4575b5050505050905090565b60005460ff161561236a576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600780546001600160a01b031990811661f002179091556008805490911661f00117905560005b8181101561266b5760008383838181106123a757fe5b905060200201356001600160a01b03166001600160a01b03161415612413576040805162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f72206164647265737300000000000000604482015290519081900360640190fd5b61243783838381811061242257fe5b905060200201356001600160a01b03166116f8565b61248657600383838381811061244957fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b6124aa83838381811061249557fe5b905060200201356001600160a01b0316612268565b6124f95760048383838181106124bc57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b600060018185858581811061250a57fe5b6001600160a01b03602091820293909301358316845283019390935260409091016000205416919091141590506125bf5782828281811061254757fe5b905060200201356001600160a01b03166001600085858581811061256757fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6000600160008585858181106125d157fe5b602090810292909201356001600160a01b031683525081019190915260400160002054600160a01b900460ff16600481111561260957fe5b14156126635760026001600085858581811061262157fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff60a01b1916600160a01b83600481111561265d57fe5b02179055505b600101612391565b50506000805460ff1916600117905550565b6000805460ff166126c4576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b6001600160a01b038c16612715576040805162461bcd60e51b8152602060048201526013602482015272496e76616c696420666565206164647265737360681b604482015290519081900360640190fd5b6128248b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c908190840183828082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b908190840183828082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a908190840183828082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a915089908190840183828082843760009201919091525061311d92505050565b61286b576040805162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9b1b934b83a34b7b760691b604482015290519081900360640190fd5b336000818152600160205260408120548190600160a01b900460ff16600481111561289257fe5b1415612990576007546040805163416259d960e11b81526001600160a01b038581166004830152915191909216916382c4b3b2916024808301926020929190829003018186803b1580156128e557600080fd5b505afa1580156128f9573d6000803e3d6000fd5b505050506040513d602081101561290f57600080fd5b5051612962576040805162461bcd60e51b815260206004820152601c60248201527f596f75206d75737420626520617574686f72697a656420666972737400000000604482015290519081900360640190fd5b506001600160a01b0381166000908152600160208190526040909120805460ff60a01b1916600160a01b1790555b6001600160a01b038281166000908152600160205260409020548116908f1614612a02578d60016000846001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6040518060a001604052808e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8f018190048102820181019092528d815291810191908e908e9081908401838280828437600092019190915250505090825250604080516020601f8d018190048102820181019092528b815291810191908c908c9081908401838280828437600092019190915250505090825250604080516020601f8b0181900481028201810190925289815291810191908a908a9081908401838280828437600092019190915250505090825250604080516020601f89018190048102820181019092528781529181019190889088908190840183828082843760009201829052509390945250506001600160a01b03851681526001602090815260409091208351805160029092019350612b62928492910190614622565b506020828101518051612b7b9260018501920190614622565b5060408201518051612b97916002840191602090910190614622565b5060608201518051612bb3916003840191602090910190614622565b5060808201518051612bcf916004840191602090910190614622565b509050508015612c29578d6001600160a01b0316826001600160a01b03167f887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22426040518082815260200191505060405180910390a3612c75565b8d6001600160a01b0316826001600160a01b03167fb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0426040518082815260200191505060405180910390a35b5060019d9c50505050505050505050505050565b6060806060806060612c996145d6565b6001600160a01b038781166000908152600160209081526040918290208251610100810190935280549384168352919290830190600160a01b900460ff166004811115612ce257fe5b6004811115612ced57fe5b815260200160018201548152602001600282016040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612da45780601f10612d7957610100808354040283529160200191612da4565b820191906000526020600020905b815481529060010190602001808311612d8757829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612e465780601f10612e1b57610100808354040283529160200191612e46565b820191906000526020600020905b815481529060010190602001808311612e2957829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f81018390048302850183019091528084529381019390830182828015612ed85780601f10612ead57610100808354040283529160200191612ed8565b820191906000526020600020905b815481529060010190602001808311612ebb57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015612f6c5780601f10612f4157610100808354040283529160200191612f6c565b820191906000526020600020905b815481529060010190602001808311612f4f57829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156130005780601f10612fd557610100808354040283529160200191613000565b820191906000526020600020905b815481529060010190602001808311612fe357829003601f168201915b5050505050815250508152602001600782015481526020016008820154815260200160098201548152602001600a820180548060200260200160405190810160405280929190818152602001828054801561308457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613066575b5050509190925250505060609081015180516020820151604083015193830151608090930151919b909a50929850909650945092505050565b60606004805480602002602001604051908101604052809291908181526020018280548015612312576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116122f4575050505050905090565b600060468651111561316f576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840dadedcd2d6cae440d8cadccee8d60531b604482015290519081900360640190fd5b610bb8855111156131c7576040805162461bcd60e51b815260206004820152601760248201527f496e76616c6964206964656e74697479206c656e677468000000000000000000604482015290519081900360640190fd5b608c84511115613217576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840eecac4e6d2e8ca40d8cadccee8d60531b604482015290519081900360640190fd5b608c83511115613265576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840cadac2d2d840d8cadccee8d60631b604482015290519081900360640190fd5b610118825111156132b6576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c8cae8c2d2d8e640d8cadccee8d60531b604482015290519081900360640190fd5b50600195945050505050565b6305f5e10081565b6000806132d76000613e4c565b915091509091565b601581565b334114613325576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b43600090815260096020908152604080832083805290915290205460ff1615613395576040805162461bcd60e51b815260206004820152601960248201527f426c6f636b20697320616c726561647920726577617264656400000000000000604482015290519081900360640190fd5b60005460ff166133db576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b4360009081526009602090815260408083208380528252808320805460ff1916600190811790915533808552925282205490913491600160a01b900460ff16600481111561342557fe5b1415613432575050613482565b61343d816000613f3b565b6040805182815242602082015281516001600160a01b038516927f7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04928290030190a250505b565b606481565b6000805460ff166134d0576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b3360006001600160a01b038416600090815260016020526040902054600160a01b900460ff16600481111561350157fe5b141561354a576040805162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b6001600160a01b038082166000908152600260209081526040808320938716835292815282822060019182905292909120825491830154909190156135c05760405162461bcd60e51b81526004018080602001828103825260238152602001806148326023913960400191505060405180910390fd5b60008111613610576040805162461bcd60e51b8152602060048201526018602482015277596f7520646f6e2774206861766520616e79207374616b6560401b604482015290519081900360640190fd5b6004546001148015613626575061362686612268565b8015613644575060018201546305f5e100906136429083614262565b105b156136805760405162461bcd60e51b815260040180806020018281038252604581526020018061477d6045913960600191505060405180910390fd5b600a82015460028401546000199091011461375d57600a8201805460001981019081106136a957fe5b9060005260206000200160009054906101000a90046001600160a01b031682600a018460020154815481106136da57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082600201546002600084600a0186600201548154811061372457fe5b60009182526020808320909101546001600160a01b0390811684528382019490945260409283018220938b168252929092529020600201555b81600a0180548061376a57fe5b600082815260209020810160001990810180546001600160a01b0319169055019055600182015461379b9082614262565b6001808401919091554390840155600060028401556005546137bd9082614262565b60055560018201546305f5e100111561386657815460ff60a01b1916600360a01b1782556137ea866142a4565b600754604080516315ea278160e01b81526001600160a01b038981166004830152915191909216916315ea27819160248083019260209291908290030181600087803b15801561383957600080fd5b505af115801561384d573d6000803e3d6000fd5b505050506040513d602081101561386357600080fd5b50505b856001600160a01b0316846001600160a01b03167f449002ae18e748d69a55f38514400d64f966492e593e32d6e9b8b24db98a0bc18342604051808381526020018281526020019250505060405180910390a350600195945050505050565b60008282018381101561391f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b60005b60045481101561397557836001600160a01b03166004828154811061394c57fe5b6000918252602090912001546001600160a01b0316141561396d5750613cc2565b60010161392b565b5060045460151115613a0857600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0385166001600160a01b031990911681179091556040805142815290517f1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a9181900360200190a2613cc2565b6000600160006004600081548110613a1c57fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060070154600160006004600081548110613a7057fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120600190810154929092019250905b600454811015613bdb57826001600060048481548110613abe57fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020600701546001600060048581548110613b1157fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154011015613bd3576001600060048381548110613b5157fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020600701546001600060048481548110613ba457fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600101540192509050805b600101613aa2565b508183850111613bec575050613cc2565b6040805142815290516001600160a01b038716917f1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a919081900360200190a260048181548110613c3857fe5b600091825260209182902001546040805142815290516001600160a01b03909216927f7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed92918290030190a28460048281548110613c9157fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505b505050565b60006001600160a01b038216600090815260016020526040902054600160a01b900460ff166004811115613cf757fe5b1480613d065750600354600110155b15613d105761191e565b6001600160a01b0381166000908152600160205260409020600701548015613d9957613d3c8183613f3b565b600654613d4990826138c5565b6006556001600160a01b038216600090815260016020526040902060080154613d7290826138c5565b6001600160a01b038316600090815260016020526040812060088101929092556007909101555b6040805182815242602082015281516001600160a01b038516927fe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235928290030190a25050565b60006001600160a01b038216600090815260016020526040902054600160a01b900460ff166004811115613e0f57fe5b1415613e1a5761191e565b6001600160a01b0381166000908152600160205260409020805460ff60a01b1916600160a21b17905561191e816142a4565b60008060005b600354811015613f355760046001600060038481548110613e6f57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b909104166004811115613eaa57fe5b14158015613edc575060038181548110613ec057fe5b6000918252602090912001546001600160a01b03858116911614155b15613f2d57613f246001600060038481548110613ef557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015484906138c5565b92506001909101905b600101613e52565b50915091565b81613f4557611897565b600080613f5183613e4c565b909250905080613f62575050611897565b600080836140d0576000613f7687856143da565b9050613f8c613f85828661441c565b8890614262565b925060005b60035481101561406557600060038281548110613faa57fe5b6000918252602090912001546001600160a01b0316905060046001600160a01b038216600090815260016020526040902054600160a01b900460ff166004811115613ff157fe5b141580156140115750876001600160a01b0316816001600160a01b031614155b1561405c576001600160a01b03811660009081526001602052604090206007015461403c90846138c5565b6001600160a01b0382166000908152600160205260409020600701559250825b50600101613f91565b5060008311801561407e57506001600160a01b03821615155b156140c6576001600160a01b0382166000908152600160205260409020600701546140a990846138c5565b6001600160a01b0383166000908152600160205260409020600701555b5050505050611897565b6000805b6003548110156141ec576000600382815481106140ed57fe5b6000918252602090912001546001600160a01b0316905060046001600160a01b038216600090815260016020526040902054600160a01b900460ff16600481111561413457fe5b141580156141545750876001600160a01b0316816001600160a01b031614155b156141e3576001600160a01b03811660009081526001602081905260408220015461418c908990614186908d9061441c565b906143da565b905061419884826138c5565b6001600160a01b038316600090815260016020526040902060070154929550935084916141c590826138c5565b6001600160a01b038316600090815260016020526040902060070155505b506001016140d4565b506141f78782614262565b925060008311801561421157506001600160a01b03821615155b15614259576001600160a01b03821660009081526001602052604090206007015461423c90846138c5565b6001600160a01b0383166000908152600160205260409020600701555b50505050505050565b600061391f83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250614475565b60005b600454811080156142ba57506004546001105b1561189757600481815481106142cc57fe5b6000918252602090912001546001600160a01b03838116911614156143d25760045460001901811461435f5760048054600019810190811061430a57fe5b600091825260209091200154600480546001600160a01b03909216918390811061433057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600480548061436a57fe5b6000828152602090819020820160001990810180546001600160a01b03191690559091019091556040805142815290516001600160a01b038516927f7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed928290030190a2611897565b6001016142a7565b600061391f83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061450c565b60008261442b57506000613922565b8282028284828161443857fe5b041461391f5760405162461bcd60e51b81526004018080602001828103825260218152602001806148116021913960400191505060405180910390fd5b600081848411156145045760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156144c95781810151838201526020016144b1565b50505050905090810190601f1680156144f65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000818361455b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156144c95781810151838201526020016144b1565b50600083858161456757fe5b0495945050505050565b8280548282559060005260206000209081019282156145c6579160200282015b828111156145c657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614591565b506145d292915061469c565b5090565b604080516101008101909152600080825260208201908152602001600081526020016146006146bb565b8152602001600081526020016000815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061466357805160ff1916838001178555614690565b82800160010185558215614690579182015b82811115614690578251825591602001919060010190614675565b506145d29291506146ea565b5b808211156145d25780546001600160a01b031916815560010161469d565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5b808211156145d257600081556001016146eb56fe596f75206d757374207761697420656e6f75676820626c6f636b7320746f20776974686472617720796f75722070726f66697473206166746572206c6174657374207769746864726177206f6620746869732076616c696461746f72596f7572207374616b696e6720686176656e277420756e6c6f636b656420796574596f752063616e277420756e7374616b652c2076616c696461746f72206c6973742077696c6c20626520656d7074792061667465722074686973206f7065726174696f6e2143616e2774207374616b65207768656e20796f752061726520756e7374616b696e6743616e2774207374616b6520746f20612076616c696461746f7220696e2061626e6f726d616c20737461747573536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77596f752061726520616c726561647920696e20756e7374616b696e6720737461747573596f7520617265206e6f742074686520666565207265636569766572206f6620746869732076616c696461746f72a2646970667358221220c698a5db956a1f1d220bfaa12f54bc449e40461e4ef4db603fbffa5fbdde573f64736f6c634300060c0033"