use AIS20220623135859


select top 10 * from xkPdaServer_good_tool where FMATERIALID = 118405 --and FUSEORGID = 1 and FISINVENTORY = 1
select top 10 * from t_BD_MaterialStock where FMATERIALID = 118405
select top 10 * from T_BD_UNIT where FMASTERID = 100007
select top 10 a.FMASTERID as FMATERIALID,                            --物料内码
       a.FNUMBER,                                --物料编码
       a1.FNAME,                                 --物料名称
       a1.FSPECIFICATION,                        --物料规格
       a2.FERPCLSID,
       case
           when a2.FERPCLSID = 1 then N'外购'
           when a2.FERPCLSID = 2 then N'自制'
           when a2.FERPCLSID = 3 then N'委外'
           else N'未定义' end as FErpClass,      --物料属性
       a2.FSUITE,                                --套件
       a2.FBASEUNITID,--物料基本单位内码
       a3.FNUMBER             as FBaseUnitNumber,--物料基本单位
       a2.FWEIGHTUNITID,                         --物料重量单位
       a2.FVOLUMEUNITID,--物料尺寸单位
       a2.FISPURCHASE,--允许采购
       a2.FISINVENTORY,--允许库存
       a2.FISSUBCONTRACT,--允许委外
       a2.FISSALE,--允许销售
       a2.FISPRODUCE,--允许生产
       a2.FISASSET,--允许资产
       a4.FISBATCHMANAGE,                        --是否启用批号管理
       a4.FISKFPERIOD,                           --是否启用保质期管理
       a5.FCHECKINCOMING,--来料检验
       a5.FCHECKPRODUCT,--产品检验,
       a5.FCHECKSTOCK,--库存检验,
       a5.FCHECKRETURN,--退货检验
       a5.FCHECKDELIVERY,--发货检验
       a41.FSTOCKID,                             --默认仓库
       a41.FNAME              as FStockName,     --默认仓库名称
       a.FUSEORGID,                              --使用组织内码
       b.FNUMBER              as FUseOrgNumber,  --使用组织编码
       c.FNAME                as FUseOrgName     --使用组织名称
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
         left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID = a.FUSEORGID
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052


select top 10 * from T_BD_MATERIAL where FMASTERID = 118405

select top 10 a.FMASTERID as FMATERIALID, a3.FUSEORGID, a.FUSEORGID
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
        left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID in (a.FUSEORGID, 0)
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052
and a.FMASTERID = 118405


select top 10 * from T_BD_UNIT where FNUMBER = 'zhi'--FUSEORGID = 0
select top 100 * from xkPdaServer_userInfo_tool
select top 100 * from xkPdaServer_keeperInfo_tool
select top 100 * from xkPdaServer_empInfo_tool
select top 100 * from xkPdaServer_customer_tool
select top 100 * from xkPdaServer_sale_tz_to_stockout_tool

select * from xkPdaServer_poorder_to_stockin_tool


SELECT DISTINCT TOP 500 [FBILLNO], [FSuppNumber], [FSuppName], [FUseOrgNumber] FROM [xkPdaServer_sltz_to_cgrk_tool] WHERE (FUseOrgNumber = '100')

select FUSERID, FNAME,*
from T_SEC_user
where FFORBIDSTATUS = 'A'

select top 10 * from t_PUR_POorder
select t0.FID,
       t0.FBILLNO,
       t3.FENTRYID,
       t3.FSEQ,
       st327.FMATERIALID,
       st327.FNAME,
       st327.FSPECIFICATION,
       iif(st327.FISBATCHMANAGE = '1', t3.FLOT_TEXT, '')    as FLOT_TEXT,
       cast(convert(float, t3.FBASEUNITQTY) as varchar(24)) as FMustQty,
       cast(convert(float, t3.FBASEUNITQTY) as varchar(24)) as SQTY,
       st327.FUseOrgNumber,
       st327.FISBATCHMANAGE,
       row_number() over (order by t0.FMODIFYDATE desc)     as idx
from t_PUR_POorder t0
         left join t_PUR_POorderEntry t3 on t0.FID = t3.FID
         left join t_PUR_POorderEntry_F t3_F on t3.FENTRYID = t3_F.FENTRYID
         left join t_PUR_POorderEntry_D t3_D on t3.FENTRYID = t3_D.FENTRYID
         left join t_PUR_POorderEntry_R t3_R on t3.FENTRYID = t3_R.FENTRYID
         join xkPdaServer_good_tool st327
              on t3.FMATERIALID = st327.FMATERIALID and t0.FPURCHASEORGID = st327.FUSEORGID
where /*t0.FPURCHASEorGID = 1
  and t3_F.FSETTLEorGID = 1
  and t3_D.FREQUIREorGID = 1
  and (t3_D.FRECEIVEorGID = 1 or (t3_D.FRECEIVEorGID = 0 and t3_D.FREQUIREorGID in (0, 1)))
  and*/ t3.FBFLOWID in ('', ' ', '182c38f9-e371-455e-9672-fad2b11b61e4',
                        '6af8ef8b-5bb8-4cdc-9972-2f71364b45d8',
                        '7127460d-b38e-4a2c-b783-72d5f0ac85b3',
                        'a6d79725-e25a-482a-a449-a867367c2b97',
                        'b27cec31-a12b-4530-8696-885e8f016280')
  and t0.FDOCUMENTSTATUS = 'C'
  and t0.FCANCELSTATUS = 'A'
  and t0.FCLOSESTATUS = 'A'
  and t3.FMRPFREEZESTATUS = 'A'
  and t3.FMRPTERMinATESTATUS = 'A'
  and t3.FMRPCLOSESTATUS = 'A'
  and t3.FCHANGEFLAG <> N'I'
  and t0.FBILLTYPEID <> N'b0677860cd16433895be5f520086b69f'
  and t0.FBILLTYPEID <> N'b8df755fd92b4c2baedef2439c29f793'
  and ABS(t3_D.FBASEDELIVERYMAXQTY) > ABS(t3_R.FBASESTOCKinQTY)
  and ABS(t3_D.FBASEDELIVERYMAXQTY) > ABS(t3_R.FBASEJOinQTY)
  and (st327.FISinVENTorY = '1' or t3.FROWTYPE = 'Service')
  and st327.FCHECKinCOMinG = '0'
  and t0.FBILLTYPEID in ('83d822ca3e374b4ab01e5dd46a0062bd', 'ba3ad5fc48d44271a048da26b615b589',
                         '0023240234df807511e308990e04cf6a')
  and t0.FOBJECTTYPEID = 'PUR_Purchaseorder'