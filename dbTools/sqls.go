package dbTools

//
//const (
//	//物料内码,物料编码,物料名称,物料规格,默认仓库id,单位number
//	GetGoods = `select top 500 a.FMATERIALID,
//							   a.FNUMBER,
//							   a.FNAME,
//							   a.FSPECIFICATION,
//							   a.FSTOCKID,
//							   a.FStockName,
//							   a.FBaseUnitNumber,
//							   a.FUSEORGID,
//							   a.FUseOrgNumber
//				from xkPdaServer_good_tool a
//				where a.FUSEORGID = 0 or a.FUseOrgNumber = '%s'`
//
//	//员工内码,员工编码,员工姓名
//	GetEmpInfo = `select top 500 a.FMATERIALID, a.FNUMBER, b.FNAME
//					from T_HR_EMPINFO a,
//						 T_HR_EMPINFO_L b,
//						 V_SCM_KEEPERORG c
//					where a.FMATERIALID = b.FID
//					  and b.FLOCALEID = 2052
//					  and c.FNUMBER = '%s'
//					  and a.FDOCUMENTSTATUS = 'C'
//					  and c.FDOCUMENTSTATUS = 'C'
//					  and c.FFORBIDSTATUS = 'A'
//					  and a.FUSEORGID in (c.FORGID, 0)`
//
//	//用户内码,用户名称
//	GetUserInfo = "select top 500 FUSERID, FNAME from T_SEC_user where FFORBIDSTATUS = 'A'"
//
//	//仓库内码,仓库编码,仓库名称
//	GetStockInfo = `select top 500 a.FMATERIALID,
//									a.FNUMBER,
//									a.FNAME,
//									a.FUSEORGID,
//									a.FUseOrgNumber
//					from xkPdaServer_stock_tool a
//					where a.FUSEORGID = 0 or a.FUseOrgNumber = '%s'`
//
//	//供应商内码,供应商编码,供应商名称
//	GetSupplierInfo = `select top 500 a.FMATERIALID,
//									  a.FNUMBER,
//							  		  a.FNAME,
//							  		  a.FSUPPLYCLASSIFY,
//									  a.FSUPPLYCLASSName,
//									  a.FORGID,
//									  a.FUseOrgNumber
//						from xkPdaServer_supplier_tool a
//						where a.FORGID = 0 or a.FUseOrgNumber = '%s'`
//
//	//客户内码,客户编码,客户名称
//	GetCustomerInfo = `select top 500 a.FMATERIALID,
//										a.FNUMBER,
//										b.FNAME,
//										a.FORGID,
//										a.FUseOrgNumber
//						from xkPdaServer_customer_tool a
//						where a.FORGID = 0 or a.FUseOrgNumber = '%s'`
//
//	//保管者编码,名字
//	GetKeeperInfo = `select a.FNUMBER, b.FNAME, a.FItemID
//						from v_itemclass_keeper a,
//							 v_itemclass_keeper_L b,
//							 V_SCM_KEEPERORG c
//						where a.fitemid = b.fitemid
//						  and b.FLOCALEID = 2052
//						  and a.fdocumentstatus = 'C'
//						  and a.FFORBIDSTATUS = 'A'
//						  and a.fmasterid = a.fitemid
//						  and c.FNUMBER = '%s'
//						  and c.FDOCUMENTSTATUS = 'C'
//						  and c.FFORBIDSTATUS = 'A'
//						  and a.FUSEORGID in (c.FORGID, 0)`
//	//组织内码,编码,名称
//	GetOrgInfo = "select a.FORGID, a.FNUMBER, b.FNAME from T_ORG_Organizations a, T_ORG_Organizations_L b where a.FORGID = b.FORGID and b.FLOCALEID = 2052 and a.FDOCUMENTSTATUS = 'C' and FFORBIDSTATUS = 'A'"
//
//	GetCGDDEntryInfo = `select a.FBILLNO,
//							   a.FID,
//							   b.FENTRYID,
//							   c.FNUMBER as FItemNumber,
//							   d.FNUMBER as FBaseUnitNumber,
//							   b.FMustQty - b.FSTOCKQTY as FMustQty
//						from t_PUR_POOrder a,
//							 T_PUR_POORDERENTRY b,
//							 T_BD_MATERIAL c,
//							 T_BD_UNIT d
//						where a.FID = b.FID
//						  and b.FMATERIALID = c.FMATERIALID
//						  and b.FUNITID = d.FMATERIALID
//						  and a.FDOCUMENTSTATUS = 'C'
//						  and a.FCLOSESTATUS = 'A'
//						  and b.FMustQty - b.FSTOCKQTY > 0
//						  and a.FBILLNO = '%s'`
//
//	GetCGDDMainInfo = `select a.FBILLNO,
//							   max(f.FNUMBER) as FSupplierNumber,
//							   max(g.FNAME) as FSupplierName
//						from t_PUR_POOrder a,
//							 T_PUR_POORDERENTRY b,
//							 V_SCM_KEEPERORG e,
//							 t_BD_Supplier f,
//							 T_BD_SUPPLIER_L g
//						where a.FID = b.FID
//						  and a.FSUPPLIERID = f.FMATERIALID
//						  and g.FSUPPLIERID = f.FMATERIALID
//						  and f.FDOCUMENTSTATUS = 'C'
//						  and a.FDOCUMENTSTATUS = 'C'
//						  and a.FCLOSESTATUS = 'A'
//						  and b.FMustQty - b.FSTOCKQTY > 0
//						  and e.FNUMBER = '%s'
//						  and e.FDOCUMENTSTATUS = 'C'
//						  and e.FFORBIDSTATUS = 'A'
//						  and f.FUSEORGID in (e.FORGID, 0)
//						  and a.FPURCHASEORGID in (e.FORGID, 0)
//						group by a.FBILLNO`
//
//	GetSCDDEntryInfo = `select a.FID, a.FBILLNO, b.FENTRYID, c.FNUMBER as FItemNumber, d.FNUMBER as FBaseUnitNumber, b.FMustQty
//						from T_PRD_MO a,
//							 T_PRD_MOENTRY b,
//							 T_BD_MATERIAL c,
//							 T_BD_UNIT d
//						where a.FID = b.FID and a.FBILLNO = '%s'
//						  and a.FDOCUMENTSTATUS = 'C' and b.FUNITID = d.FMATERIALID
//						and c.FMATERIALID = b.FMATERIALID
//						and c.FDOCUMENTSTATUS = 'C'`
//
//	GetSCDDMainInfo = `select a.FBILLNO
//						from T_PRD_MO a,
//							 V_SCM_KEEPERORG e
//						where a.FDOCUMENTSTATUS = 'C'
//						and e.FDOCUMENTSTATUS = 'C'
//						and e.FNUMBER = '%s'
//						and a.FPRDORGID in (e.FORGID, 0)`
//
//	GetSCTLMainInfo = `select 1`
//
//	GetSCTLEntryInfo = `select 1`
//
//	GetXSDDMainInfo = `select 1`
//
//	GetXSDDEntryInfo = `select 1`
//
//	GetWWDDMainInfo = `select 1`
//
//	GetWWDDEntryInfo = `select 1`
//
//	GetWWTLMainInfo = `select 1`
//
//	GetWWTLEntryInfo = `select 1`
//)
