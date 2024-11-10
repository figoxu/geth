

分层确定性钱包（Hierarchical Deterministic Wallet，简称HD Wallet）是一种用于加密货币管理的钱包结构，能够根据一个主密钥或种子生成一系列密钥的层次结构。简单来说，它可以在一个钱包中生成多个账户，每个账户下还能有多个地址。具体来说，分层确定性钱包有以下几个特点：

1. **层次结构**：HD钱包的层次结构有明确的父子关系，账户和子地址都按层次方式管理。主密钥（seed）可以通过不同的路径生成不同的子密钥，形成了一个「树状结构」。

2. **确定性**：通过一个种子可以确定性地生成整个钱包的密钥结构。这意味着只需备份一个种子词（通常是12或24个助记词）就可以恢复整个钱包的密钥和账户。

3. **安全性和隐私性**：HD钱包通常使用一种算法（比如BIP32）来生成公私钥对。在需要生成新地址时，不需要生成新的私钥，而是从现有的主密钥派生出新地址，提高了安全性，同时用户可以随时更换地址以提升隐私性。

4. **简化管理**：因为钱包的所有地址和账户都是从一个种子生成的，所以备份和恢复整个钱包的过程变得非常简单。这对于用户的资产管理来说非常方便，因为只要保存好种子词，就可以随时恢复钱包。

在实际使用中，HD钱包大大简化了多账户、多地址管理的复杂度，提高了用户的安全性和隐私性，是目前加密货币钱包的主流技术之一。


https://github.com/miguelmota/go-ethereum-hdwallet
